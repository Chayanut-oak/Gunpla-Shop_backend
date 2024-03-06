package dynamoDB

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
		var order entity.Order
		err := attributevalue.UnmarshalMap(item, &order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (repo *OrderRepository) AddOrder(order restModel.OrderRestModel, orderId string) (*restModel.OrderRestModel, error) {
	item, err := attributevalue.MarshalMap(order)
	currentTime := time.Now().Format(time.DateTime)
	item["OrderId"] = &types.AttributeValueMemberS{Value: orderId}
	item["OrderDate"] = &types.AttributeValueMemberS{Value: currentTime} // Replace "your_order_date_here" with the actual order date
	item["Status"] = &types.AttributeValueMemberS{Value: "Pending"}      // Replace "your_status_here" with the actual status
	delete(item, "Token")
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
func (repo *OrderRepository) UpdateOrderStock(order restModel.OrderRestModel) (string, error) {
	// Start a transaction
	transaction := make([]types.TransactWriteItem, 0, len(order.Cart))
	for _, item := range order.Cart {
		// fmt.Println(item)
		// Marshal the key for the update operation
		key, err := attributevalue.MarshalMap(map[string]string{
			"ProductId": item.ProductId,
		})
		if err != nil {
			return "", fmt.Errorf("failed to marshal key: %v", err)
		}

		// Define the update operation
		update := &types.Update{
			TableName: aws.String(item.Type + "s"),
			Key:       key,
			// Update stock with conditional expression to eznsure it does not become negative
			UpdateExpression:          aws.String("SET Stock = if_not_exists(Stock, :initial) - :quantity"),
			ExpressionAttributeValues: map[string]types.AttributeValue{":quantity": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", item.Quantity)}, ":initial": &types.AttributeValueMemberN{Value: "0"}},
			ConditionExpression:       aws.String("attribute_exists(Stock) and Stock >= :quantity"),
		}

		// Add the update operation to the transaction
		transaction = append(transaction, types.TransactWriteItem{Update: update})
	}

	// If no valid updates were added to the transaction, return nil
	if len(transaction) == 0 {
		return "", nil
	}

	// Execute the transaction
	_, err := repo.Client.TransactWriteItems(context.TODO(), &dynamodb.TransactWriteItemsInput{
		TransactItems: transaction,
	})
	if err != nil {
		log.Printf("Failed to execute transaction: %v", err)
		return "", fmt.Errorf("transaction failed: %v", err)
	}

	return "Success", nil
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
	if order.Status == "Cancel" {
		transaction := make([]types.TransactWriteItem, 0, len(order.Cart))
		for _, item := range order.Cart {
			// fmt.Println(item)
			// Marshal the key for the update operation
			key, err := attributevalue.MarshalMap(map[string]string{
				"ProductId": item.ProductId,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to marshal key: %v", err)
			}

			// Define the update operation
			update := &types.Update{
				TableName: aws.String(item.Type + "s"),
				Key:       key,
				// Update stock with conditional expression to eznsure it does not become negative
				UpdateExpression:          aws.String("SET Stock = if_not_exists(Stock, :initial) + :quantity"),
				ExpressionAttributeValues: map[string]types.AttributeValue{":quantity": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", item.Quantity)}, ":initial": &types.AttributeValueMemberN{Value: "0"}},
				ConditionExpression:       aws.String("attribute_exists(Stock) and Stock >= :quantity"),
			}

			// Add the update operation to the transaction
			transaction = append(transaction, types.TransactWriteItem{Update: update})
		}

		// If no valid updates were added to the transaction, return nil
		if len(transaction) == 0 {
			return nil, err
		}

		// Execute the transaction
		_, err := repo.Client.TransactWriteItems(context.TODO(), &dynamodb.TransactWriteItemsInput{
			TransactItems: transaction,
		})
		if err != nil {
			log.Printf("Failed to execute transaction: %v", err)
			return nil, fmt.Errorf("transaction failed: %v", err)
		}
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
