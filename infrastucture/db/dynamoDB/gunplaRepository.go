package dynamoDB

import (
	"context"
	"fmt"
	"log"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type GunplaRepository struct {
	Client *dynamodb.Client
}

func CreateGunplaRepository(client *dynamodb.Client) *GunplaRepository {
	return &GunplaRepository{Client: client}
}

func (repo *GunplaRepository) GetAllGunplas() ([]*entity.Gunpla, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String("Gunplas"),
	}
	result, err := repo.Client.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan DynamoDB table: %v", err)
	}
	var gunplas []*entity.Gunpla
	for _, item := range result.Items {
		fmt.Println(item)
		var gunpla entity.Gunpla
		err := attributevalue.UnmarshalMap(item, &gunpla)
		if err != nil {
			return nil, err
		}
		gunplas = append(gunplas, &gunpla)
		fmt.Println(gunplas)
	}
	fmt.Println(gunplas)
	return gunplas, nil
}

func (repo *GunplaRepository) AddGunpla(gunpla entity.NewGunpla) (*entity.NewGunpla, error) {
	item, err := attributevalue.MarshalMap(gunpla)
	item["GunplaId"] = &types.AttributeValueMemberS{Value: uuid.NewString()}
	fmt.Print(item)
	if err != nil {
		return nil, err
	}

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Gunplas"),
		Item:      item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
		return nil, err
	}
	return &gunpla, nil
}

// func (repo *GunplaRepository) UpdateGunpla(gunpla entity.Gunpla) (*entity.Gunpla, error) {
// 	input := &dynamodb.UpdateItemInput{
// 		TableName: aws.String("Gunplas"),
// 		Key: map[string]types.AttributeValue{
// 			"gunplaId": &types.AttributeValueMemberS{Value: gunpla.gunplaId}, // Assuming ID is the primary key
// 		},
// 		UpdateExpression: aws.String("SET #attrName = :attrValue"),
// 		ExpressionAttributeNames: map[string]string{
// 			"#attrName": "AttributeName", // Replace AttributeName with the actual attribute name you want to update
// 		},
// 		ExpressionAttributeValues: map[string]types.AttributeValue{
// 			":attrValue": &types.AttributeValueMemberS{Value: gunpla.NewValue}, // Replace NewValue with the new value you want to set
// 		},
// 		ReturnValues: types.ReturnValueUpdatedNew,
// 	}

// 	// Perform the update operation
// 	result, err := repo.Client.UpdateItem(context.Background(), input)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to update item: %v", err)
// 	}

// 	// Parse and return the updated item
// 	updatedGunpla := &entity.Gunpla{
// 		ID: gunpla.ID,
// 		// Set other attributes if needed
// 	}
// 	return updatedGunpla, nil
// }
