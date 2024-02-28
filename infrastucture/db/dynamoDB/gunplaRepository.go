package dynamoDB

import (
	"context"
	"fmt"
	"log"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
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

func (repo *GunplaRepository) UpdateGunpla(gunpla entity.Gunpla) (*entity.Gunpla, error) {
	key, err := attributevalue.MarshalMap(map[string]string{
		"GunplaId": gunpla.GunplaId,
	})
	if err != nil {
		return nil, err
	}
	update := expression.Set(expression.Name("Images"), expression.Value(gunpla.Images))
	update.Set(expression.Name("Name"), expression.Value(gunpla.Name))
	update.Set(expression.Name("Type"), expression.Value(gunpla.Type))
	update.Set(expression.Name("Series"), expression.Value(gunpla.Series))
	update.Set(expression.Name("Scale"), expression.Value(gunpla.Scale))
	update.Set(expression.Name("Grade"), expression.Value(gunpla.Grade))
	update.Set(expression.Name("Price"), expression.Value(gunpla.Price))
	update.Set(expression.Name("Stock"), expression.Value(gunpla.Stock))
	update.Set(expression.Name("Description"), expression.Value(gunpla.Description))

	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		log.Printf("Couldn't build expression for update. Here's why: %v\n", err)
		return nil, err
	}

	_, err = repo.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String("Gunplas"),
		Key:                       key,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	})
	if err != nil {
		log.Printf("Couldn't update item in table. Here's why: %v\n", err)
		return nil, err
	}

	return &gunpla, err
}

func (repo *GunplaRepository) DeleteGunpla(gunpla entity.Gunpla) error {
	key, err := attributevalue.MarshalMap(map[string]string{
		"GunplaId": gunpla.GunplaId,
	})
	if err != nil {
		return err
	}
	_, err = repo.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String("Gunplas"), Key: key,
	})

	if err != nil {
		log.Printf("Couldn't delete item in table. Here's why: %v\n", err)
		return err
	}
	return err
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
