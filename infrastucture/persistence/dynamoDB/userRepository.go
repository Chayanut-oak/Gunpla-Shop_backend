package dynamoDB

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	Client *dynamodb.Client
}

func CreateUserRepository(client *dynamodb.Client) *UserRepository {
	return &UserRepository{Client: client}
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (repo *UserRepository) NewUser(user restModel.UserRestModel) (string, error) {
	hashedPassword, err := hashPassword(user.Password)

	if err != nil {
		return "", err
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	isValid := emailRegex.MatchString(user.Email)

	if !isValid {
		errMsg := fmt.Sprintf("Invalid email address: %s", user.Email)
		return errMsg, fmt.Errorf(errMsg)
	}
	user.Password = hashedPassword
	user.Role = "customer"
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return "", err
	}

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName:           aws.String("Users"),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(Email)"),
	})
	if err != nil {
		fmt.Printf("Couldn't Create User to table. Here's why: %v\n", err)
		return "", err
	}

	return user.Email, nil
}
func (repo *UserRepository) AuthenticateUser(email, password string) (bool, error) {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	isValid := emailRegex.MatchString(email)

	if !isValid {
		errMsg := fmt.Sprintf("Invalid email address: %s", email)
		return false, fmt.Errorf(errMsg)
	}

	result, err := repo.GetUserByEmail(email)

	if err != nil {
		return false, err
	}

	if result == nil {
		return false, errors.New("user not found")
	}

	hashedPassword := result.Password
	if hashedPassword == "" {
		return false, errors.New("password not found in user record")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Users"),
		Key: map[string]types.AttributeValue{
			"Email": &types.AttributeValueMemberS{Value: email},
		},
	}
	result, err := repo.Client.GetItem(context.TODO(), input)
	fmt.Print(result)
	if err != nil {
		return nil, errors.New("error getting user by Email")
	}

	var user entity.User
	err = attributevalue.UnmarshalMap(result.Item, &user)
	if err != nil {
		fmt.Println("Error unmarshaling DynamoDB result:", err)
		return nil, fmt.Errorf("error unmarshaling DynamoDB result: %w", err)
	}

	return &user, nil
}
