package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-3"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	client := sts.NewFromConfig(cfg)
	client.GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("Call to STS get-caller-identity failed, %v", err)
	}
	presignedClient := sts.NewPresignClient(client)
	identitySigned, err := presignedClient.PresignGetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("Failed to generate pre-signed URL for STS get caller identity, %v", err)
	}
	fmt.Println(identitySigned.URL)
}
