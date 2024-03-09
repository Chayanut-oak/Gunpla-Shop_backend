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
	godotenv.Load()
	accessKey := os.Getenv("ACCESSKEYID")
	secretAccessKey := os.Getenv("SERCETCCESSKEY")
	fmt.Println("accessKey ", accessKey)
	fmt.Println("secretAccessKey ", secretAccessKey)
	fmt.Println(accessKey, secretAccessKey)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("localhost"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKey,
				SecretAccessKey: secretAccessKey,
			},
		}),
	)
	if err != nil {
		panic(err)
	}
	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDBClient{Client: client}, nil
}
