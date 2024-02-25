package infrastructure

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb/expression"
    "github.com/Chayanut-oak/Gunpla-Shop_backend/domain"
)

// DynamoDBProductRepository implements the ProductRepository interface using DynamoDB
type DynamoDBProductRepository struct {
    TableName string
    Client    *dynamodb.DynamoDB
}

// NewDynamoDBProductRepository creates a new instance of DynamoDBProductRepository
func NewDynamoDBProductRepository(tableName, region string) *DynamoDBProductRepository {
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String(region),
    }))
    return &DynamoDBProductRepository{
        TableName: tableName,
        Client:    dynamodb.New(sess),
    }
}

// GetProductByID retrieves a product from DynamoDB by ID
func (repo *DynamoDBProductRepository) GetProductByID(id int) (*domain.Product, error) {
    key := map[string]*dynamodb.AttributeValue{
        "ID": {
            N: aws.String(string(id)),
        },
    }

    result, err := repo.Client.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String(repo.TableName),
        Key:       key,
    })
    if err != nil {
        return nil, err
    }

    if result.Item == nil {
        return nil, nil // Product not found
    }

    var product domain.Product
    err = dynamodbattribute.UnmarshalMap(result.Item, &product)
    if err != nil {
        return nil, err
    }

    return &product, nil
}
