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
	"github.com/google/uuid"
)

type ForumRepository struct {
	Client *dynamodb.Client
}

func CreateForumRepository(client *dynamodb.Client) *ForumRepository {
	return &ForumRepository{Client: client}
}

func (repo *ForumRepository) GetAllForums() ([]*entity.Forum, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String("Forums"),
	}
	result, err := repo.Client.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan DynamoDB table: %v", err)
	}
	var forums []*entity.Forum
	for _, item := range result.Items {
		var forum entity.Forum
		err := attributevalue.UnmarshalMap(item, &forum)
		if err != nil {
			return nil, err
		}
		forums = append(forums, &forum)
	}
	return forums, nil
}
func (repo *ForumRepository) GetForum(forumId string) (*entity.Forum, error) {
	// Define the query input parameters
	input := &dynamodb.QueryInput{
		TableName:              aws.String("Forums"), // Specify the table name
		KeyConditionExpression: aws.String("ForumId = :fid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":fid": &types.AttributeValueMemberS{Value: forumId},
		},
	}

	// Execute the query
	result, err := repo.Client.Query(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to query DynamoDB table: %v", err)
	}

	// Check if the query result has any items
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("forum not found")
	}

	// Unmarshal the query result into *entity.Forum
	var forum entity.Forum
	err = attributevalue.UnmarshalMap(result.Items[0], &forum) // Assuming ForumId is unique, so we only expect one item
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal DynamoDB item: %v", err)
	}

	return &forum, nil
}
func (repo *ForumRepository) AddForum(forum restModel.ForumRestModel) (*restModel.ForumRestModel, error) {
	item, err := attributevalue.MarshalMap(forum)
	currentTime := time.Now().Format(time.DateTime)
	item["ForumId"] = &types.AttributeValueMemberS{Value: uuid.NewString()}
	item["CreatedAt"] = &types.AttributeValueMemberS{Value: currentTime} // Replace "your_forum_date_here" with the actual forum date

	if err != nil {
		return nil, err
	}
	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Forums"),
		Item:      item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
		return nil, err
	}
	return &forum, nil
}

func (repo *ForumRepository) UpdateForum(forum entity.Forum) (*entity.Forum, error) {
	item, err := attributevalue.MarshalMap(forum)
	fmt.Print(item)
	if err != nil {
		return nil, err
	}

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Forums"),
		Item:      item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
		return nil, err
	}
	return &forum, nil
}

func (repo *ForumRepository) DeleteForum(ForumId string) error {

	key, err := attributevalue.MarshalMap(map[string]string{
		"ForumId": ForumId,
	})

	if err != nil {
		return err
	}
	_, err = repo.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String("Forums"), Key: key,
	})

	if err != nil {
		log.Printf("Couldn't delete item in table. Here's why: %v\n", err)
		return err
	}
	return err
}

func (repo *ForumRepository) GetAllCommentsInForum(forumId string) ([]*entity.Comment, error) {
	fmt.Println("forumId", forumId)
	input := &dynamodb.QueryInput{
		TableName:              aws.String("Comments"),
		KeyConditionExpression: aws.String("#sk = :skval"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":skval": &types.AttributeValueMemberS{Value: forumId},
		},
		ExpressionAttributeNames: map[string]string{
			"#sk": "ForumId",
		},
	}
	fmt.Println("input", input)

	result, err := repo.Client.Query(context.TODO(), input)
	if err != nil {
		fmt.Println("err", err)
		return nil, fmt.Errorf("failed to query DynamoDB table: %v", err)
	}
	fmt.Println("result", result)

	var comments []*entity.Comment
	for _, item := range result.Items {
		var comment entity.Comment
		err := attributevalue.UnmarshalMap(item, &comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return comments, nil

}

func (repo *ForumRepository) AddComment(comment restModel.CommentRestModel) (*restModel.CommentRestModel, error) {
	item, err := attributevalue.MarshalMap(comment)
	currentTime := time.Now().Format(time.DateTime)
	item["CommentId"] = &types.AttributeValueMemberS{Value: uuid.NewString()}
	item["CreatedAt"] = &types.AttributeValueMemberS{Value: currentTime} // Replace "your_forum_date_here" with the actual forum date

	if err != nil {
		return nil, err
	}
	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Comments"),
		Item:      item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
		return nil, err
	}
	return &comment, nil
}

func (repo *ForumRepository) UpdateComment(comment entity.Comment) (*entity.Comment, error) {
	item, err := attributevalue.MarshalMap(comment)
	fmt.Print(item)
	if err != nil {
		return nil, err
	}

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Comments"),
		Item:      item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
		return nil, err
	}
	return &comment, nil
}

func (repo *ForumRepository) DeleteComment(CommentId string) error {

	key, err := attributevalue.MarshalMap(map[string]string{
		"CommentId": CommentId,
	})

	if err != nil {
		return err
	}
	_, err = repo.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String("Comments"), Key: key,
	})

	if err != nil {
		log.Printf("Couldn't delete item in table. Here's why: %v\n", err)
		return err
	}
	return err
}
