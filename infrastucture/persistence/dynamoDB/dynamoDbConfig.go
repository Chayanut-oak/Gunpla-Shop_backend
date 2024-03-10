package dynamoDB

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/joho/godotenv"
)

type DynamoDBClient struct {
	Client *dynamodb.Client
}

func CreateDynamoDBClient() (*DynamoDBClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion("us-east-1"),
    )
    if err != nil {
        return nil, fmt.Errorf("failed to load AWS config: %v", err)
    }

    client := dynamodb.NewFromConfig(cfg)
    return &DynamoDBClient{Client: client}, nil
}
