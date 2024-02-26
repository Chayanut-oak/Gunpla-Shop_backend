package dynamoDB

import (
	"context"
	"fmt"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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
