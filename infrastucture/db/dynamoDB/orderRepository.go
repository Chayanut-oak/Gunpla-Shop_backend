package dynamoDB

import (
	"context"
	"fmt"
	"log"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type OrderRepository struct {
	Client *dynamodb.Client
}

func CreateOrderRepository(client *dynamodb.Client) *OrderRepository {
	return &OrderRepository{Client: client}
}

func (repo *OrderRepository) GetAllOrders() ([]*entity.Order, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String("Orders"),
	}
	result, err := repo.Client.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan DynamoDB table: %v", err)
	}
	var orders []*entity.Order
	for _, item := range result.Items {
		fmt.Println(item)
		var order entity.Order
		err := attributevalue.UnmarshalMap(item, &order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
		fmt.Println(orders)
	}
	fmt.Println(orders)
	return orders, nil
}

func (repo *OrderRepository) AddOrder(order restModel.OrderRestModal) (*restModel.OrderRestModal, error) {
	item, err := attributevalue.MarshalMap(order)
	item["OrderId"] = &types.AttributeValueMemberS{Value: uuid.NewString()}
	fmt.Print(item)
	if err != nil {
		return nil, err
	}
	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Orders"),
		Item:      item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
		return nil, err
	}
	return &order, nil
}

// func (repo *OrderRepository) UpdateOrder(order entity.Order) (*entity.Order, error) {
// 	key, err := attributevalue.MarshalMap(map[string]string{
// 		"OrderId": order.OrderId,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	update := expression.Set(expression.Name("Images"), expression.Value(order.Images))
// 	update.Set(expression.Name("Name"), expression.Value(order.Name))
// 	update.Set(expression.Name("Type"), expression.Value(order.Type))
// 	update.Set(expression.Name("Series"), expression.Value(order.Series))
// 	update.Set(expression.Name("Scale"), expression.Value(order.Scale))
// 	update.Set(expression.Name("Grade"), expression.Value(order.Grade))
// 	update.Set(expression.Name("Price"), expression.Value(order.Price))
// 	update.Set(expression.Name("Stock"), expression.Value(order.Stock))
// 	update.Set(expression.Name("Description"), expression.Value(order.Description))

// 	expr, err := expression.NewBuilder().WithUpdate(update).Build()
// 	if err != nil {
// 		log.Printf("Couldn't build expression for update. Here's why: %v\n", err)
// 		return nil, err
// 	}

// 	_, err = repo.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
// 		TableName:                 aws.String("Orders"),
// 		Key:                       key,
// 		ExpressionAttributeNames:  expr.Names(),
// 		ExpressionAttributeValues: expr.Values(),
// 		UpdateExpression:          expr.Update(),
// 	})
// 	if err != nil {
// 		log.Printf("Couldn't update item in table. Here's why: %v\n", err)
// 		return nil, err
// 	}

// 	return &order, err
// }

func (repo *OrderRepository) UpdateOrder(order entity.Order) (*entity.Order, error) {
	item, err := attributevalue.MarshalMap(order)
	fmt.Print(item)
	if err != nil {
		return nil, err
	}

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Orders"),
		Item:      item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
		return nil, err
	}
	return &order, nil
}

func (repo *OrderRepository) DeleteOrder(OrderId string) error {

	key, err := attributevalue.MarshalMap(map[string]string{
		"OrderId": OrderId,
	})

	if err != nil {
		return err
	}
	_, err = repo.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String("Orders"), Key: key,
	})

	if err != nil {
		log.Printf("Couldn't delete item in table. Here's why: %v\n", err)
		return err
	}
	return err
}

// func (repo *OrderRepository) UpdateOrder(order entity.Order) (*entity.Order, error) {
// 	input := &dynamodb.UpdateItemInput{
// 		TableName: aws.String("Orders"),
// 		Key: map[string]types.AttributeValue{
// 			"orderId": &types.AttributeValueMemberS{Value: order.orderId}, // Assuming ID is the primary key
// 		},
// 		UpdateExpression: aws.String("SET #attrName = :attrValue"),
// 		ExpressionAttributeNames: map[string]string{
// 			"#attrName": "AttributeName", // Replace AttributeName with the actual attribute name you want to update
// 		},
// 		ExpressionAttributeValues: map[string]types.AttributeValue{
// 			":attrValue": &types.AttributeValueMemberS{Value: order.NewValue}, // Replace NewValue with the new value you want to set
// 		},
// 		ReturnValues: types.ReturnValueUpdatedNew,
// 	}

// 	// Perform the update operation
// 	result, err := repo.Client.UpdateItem(context.Background(), input)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to update item: %v", err)
// 	}

// 	// Parse and return the updated item
// 	updatedOrder := &entity.Order{
// 		ID: order.ID,
// 		// Set other attributes if needed
// 	}
// 	return updatedOrder, nil
// }