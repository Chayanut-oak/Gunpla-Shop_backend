package dynamoDB

import (
	"context"
	"fmt"
	"log"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"

	// "github.com/Chayanut-oak/Gunpla-Shop_backend/domain/valueObject"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type ToolRepository struct {
	Client *dynamodb.Client
}

func CreateToolRepository(client *dynamodb.Client) *ToolRepository {
	return &ToolRepository{Client: client}
}

func (repo *ToolRepository) GetAllTools() ([]*entity.Tool, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String("Tools"),
	}
	result, err := repo.Client.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan DynamoDB table: %v", err)
	}
	var tools []*entity.Tool
	for _, item := range result.Items {
		fmt.Println(item)
		var tool entity.Tool
		err := attributevalue.UnmarshalMap(item, &tool)
		if err != nil {
			return nil, err
		}
		tools = append(tools, &tool)
		fmt.Println(tools)
	}
	fmt.Println(tools)
	return tools, nil
}

func (repo *ToolRepository) AddTool(tool restModel.ToolRestModal) (*restModel.ToolRestModal, error) {
	fmt.Print(tool)
	item, err := attributevalue.MarshalMap(tool)
	item["ProductId"] = &types.AttributeValueMemberS{Value: uuid.NewString()}
	fmt.Print(item)
	if err != nil {
		return nil, err
	}

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Tools"),
		Item:      item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
		return nil, err
	}
	return &tool, nil
}

func (repo *ToolRepository) UpdateTool(tool entity.Tool) (*entity.Tool, error) {
	item, err := attributevalue.MarshalMap(tool)
	fmt.Print(item)
	if err != nil {
		return nil, err
	}

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Tools"),
		Item:      item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
		return nil, err
	}
	return &tool, nil
}

// func (repo *ToolRepository) UpdateToolStock(tool []valueObject.Product) (string, error) {
// 	// Start a transaction
// 	transaction := make([]types.TransactWriteItem, 0, len(tool))

// 	for _, item := range tool {
// 		// Marshal the key for the update operation
// 		key, err := attributevalue.MarshalMap(map[string]string{
// 			"ProductId": item.ProductId,
// 		})
// 		if err != nil {
// 			return "", fmt.Errorf("failed to marshal key: %v", err)
// 		}

// 		// Define the update operation
// 		update := &types.Update{
// 			TableName: aws.String("Tools"),
// 			Key:       key,
// 			// Update stock with conditional expression to ensure it does not become negative
// 			UpdateExpression:          aws.String("SET Stock = if_not_exists(Stock, :initial) - :quantity"),
// 			ExpressionAttributeValues: map[string]types.AttributeValue{":quantity": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", item.Quantity)}, ":initial": &types.AttributeValueMemberN{Value: "0"}},
// 			ConditionExpression:       aws.String("attribute_exists(Stock) and Stock >= :quantity"),
// 		}

// 		// Add the update operation to the transaction
// 		transaction = append(transaction, types.TransactWriteItem{Update: update})
// 	}

// 	// If no valid updates were added to the transaction, return nil
// 	if len(transaction) == 0 {
// 		return "", nil
// 	}

// 	// Execute the transaction
// 	_, err := repo.Client.TransactWriteItems(context.TODO(), &dynamodb.TransactWriteItemsInput{
// 		TransactItems: transaction,
// 	})
// 	if err != nil {
// 		log.Printf("Failed to execute transaction: %v", err)
// 		return "", fmt.Errorf("transaction failed: %v", err)
// 	}

// 	return "Success", nil
// }

func (repo *ToolRepository) DeleteTool(ProductId string) error {

	key, err := attributevalue.MarshalMap(map[string]string{
		"ProductId": ProductId,
	})

	if err != nil {
		return err
	}
	_, err = repo.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String("Tools"), Key: key,
	})

	if err != nil {
		log.Printf("Couldn't delete item in table. Here's why: %v\n", err)
		return err
	}
	return err
}
