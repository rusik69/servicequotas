package main

import (
	"flag"
	"log"

	"github.com/rusik69/servicequotas/pkg/aws"
	"github.com/rusik69/servicequotas/pkg/config"
)

func main() {
	region := flag.String("region", "", "AWS region")
	configFileName := flag.String("config", "", "quotas config file")
	flag.Parse()
	if *region == "" {
		log.Fatal("region is required")
	}
	if *configFileName == "" {
		log.Fatal("config is required")
	}
	cfg, err := config.Parse(*configFileName)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	log.Printf("config parsed")
	awsSession, err := aws.CreateSession(*region)
	if err != nil {
		log.Fatalf("failed to create AWS session: %v", err)
	}
	log.Println("AWS session created")
	log.Printf("adjusting quotas")
	requests, err := aws.AdjustQuotas(awsSession, cfg)
	if err != nil {
		log.Fatalf("failed to adjust quotas: %v", err)
	}
	aws.WaitForRequests(awsSession, requests)
	log.Printf("quotas adjusted")
}
