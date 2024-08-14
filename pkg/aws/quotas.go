package aws

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	sqtypes "github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/rusik69/servicequotas/pkg/types"
)

// GetServiceQuota retrieves the service quota for a specific service and quota code
func GetServiceQuota(awsCfg aws.Config, serviceCode, quotaCode string) (*servicequotas.GetServiceQuotaOutput, error) {
	svc := servicequotas.NewFromConfig(awsCfg)

	input := &servicequotas.GetServiceQuotaInput{
		ServiceCode: &serviceCode,
		QuotaCode:   &quotaCode,
	}

	result, err := svc.GetServiceQuota(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// AdjustQuotas adjusts the quotas
func AdjustQuotas(awsCfg aws.Config, cfg types.QuotasConfig) ([]string, error) {
	// requests is a list of request ids
	var requests []string
	session := servicequotas.NewFromConfig(awsCfg)
	for _, q := range cfg.Quotas {
		log.Printf("checking quota %s", q.QuotaName)
		if q.Adjustable {
			log.Printf("getting current quota %s", q.QuotaName)
			currentQuota, err := GetServiceQuota(awsCfg, q.ServiceCode, q.QuotaCode)
			if err != nil {
				log.Printf("failed to get current quota %s value: %v", q.QuotaName, err)
				continue
			}
			if int(*currentQuota.Quota.Value) < int(q.Value) {
				log.Printf("requesting increase for quota %s from %f to %f", q.QuotaName, *currentQuota.Quota.Value, q.Value)
				input := &servicequotas.RequestServiceQuotaIncreaseInput{
					QuotaCode:    &q.QuotaCode,
					ServiceCode:  &q.ServiceCode,
					DesiredValue: &q.Value,
				}
				result, err := session.RequestServiceQuotaIncrease(context.TODO(), input)
				if err != nil {
					log.Printf("failed to request increase for quota %s: %v", q.QuotaName, err)
					continue
				}
				requestID := result.RequestedQuota.CaseId
				if requestID == nil {
					log.Printf("failed to get request id for quota increase request %s", q.QuotaName)
					continue
				}
				log.Printf("quota increase request %s submitted with request id: %s", q.QuotaName, *requestID)
				requests = append(requests, *requestID)
			}
		} else {
			log.Printf("quota %s is not adjustable, skipping", q.QuotaName)
		}
	}
	return requests, nil
}

// WaitForRequests waits for the quota increase requests to complete
func WaitForRequests(awsCfg aws.Config, requests []string) {
	session := servicequotas.NewFromConfig(awsCfg)
	var wg sync.WaitGroup
	for _, requestID := range requests {
		wg.Add(1)
		go func(requestID string) {
			defer wg.Done()
			for {
				statusInput := &servicequotas.GetRequestedServiceQuotaChangeInput{
					RequestId: &requestID,
				}
				statusResult, err := session.GetRequestedServiceQuotaChange(context.TODO(), statusInput)
				if err != nil {
					log.Printf("failed to get status for quota increase request %s: %v", requestID, err)
					break
				}
				if statusResult.RequestedQuota.Status == sqtypes.RequestStatusCaseClosed ||
					statusResult.RequestedQuota.Status == sqtypes.RequestStatusNotApproved {
					log.Printf("quota increase request %s completed with status: %s", requestID, statusResult.RequestedQuota.Status)
					break
				}
				log.Printf("waiting for quota increase request %s to complete, current status: %s", requestID, statusResult.RequestedQuota.Status)
				time.Sleep(30 * time.Second) // Wait for 30 seconds before checking again
			}
		}(requestID)
	}
	wg.Wait()
}
