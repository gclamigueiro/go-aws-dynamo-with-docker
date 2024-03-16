package daoImpl

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gclamigueiro/go-aws-dynamo-with-docker/internal/dao"
	"github.com/gclamigueiro/go-aws-dynamo-with-docker/internal/dynamo"
	"github.com/gclamigueiro/go-aws-dynamo-with-docker/internal/models"
)

type IndexAttributes struct {
	PK string `dynamodbav:"pk"`
	SK string `dynamodbav:"sk"`
}

type UserItem struct {
	models.User
	IndexAttributes
}

// GetKey returns the composite primary key of the user in a format that can be sent to DynamoDB.
func (user UserItem) GetKey() map[string]types.AttributeValue {
	pk, err := attributevalue.Marshal(user.PK)
	if err != nil {
		panic(err)
	}
	sk, err := attributevalue.Marshal(user.SK)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"pk": pk, "sk": sk}
}

func getPK(username string) string {
	return "USER#" + username
}

type userDao struct {
	baseDao
}

func NewUserDao(client *dynamodb.Client) dao.UserDao {
	return userDao{
		baseDao: baseDao{
			client: client,
		},
	}
}

func (dao userDao) GetUser(ctx context.Context, username string) (models.User, error) {

	user := UserItem{
		IndexAttributes: IndexAttributes{
			PK: getPK(username),
			SK: getPK(username),
		},
	}
	response, err := dao.client.GetItem(ctx, &dynamodb.GetItemInput{
		Key: user.GetKey(), TableName: aws.String(dynamo.MainTableName),
	})
	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", username, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &user)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}

	return user.User, nil
}

func (dao userDao) AddUser(ctx context.Context, user models.User) error {

	userItem := UserItem{
		User: user,
		IndexAttributes: IndexAttributes{
			PK: getPK(user.Username),
			SK: getPK(user.Username),
		},
	}

	item, err := attributevalue.MarshalMap(userItem)
	if err != nil {
		panic(err)
	}
	_, err = dao.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(dynamo.MainTableName), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return err
}
