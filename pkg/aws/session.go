package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// CreateSession creates a new AWS session
func CreateSession(region string) (aws.Config, error) {
	// Load the default configuration
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		return aws.Config{}, err
	}
	return awsCfg, nil
}
