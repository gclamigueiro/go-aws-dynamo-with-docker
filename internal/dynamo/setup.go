package dynamo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

type DynamoConfigurator struct {
	client *dynamodb.Client
}

func NewDynamoConfigurator(c *dynamodb.Client) DynamoConfigurator {
	return DynamoConfigurator{
		client: c,
	}
}

func (setup DynamoConfigurator) SetupDatabase() {
	exist, err := setup.TableExists(MainTableName)
	if err != nil {
		panic(err)
	}
	if !exist {
		setup.createMainTable()
	}
}

func (setup DynamoConfigurator) TableExists(tableName string) (bool, error) {
	exists := true
	_, err := setup.client.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(tableName)},
	)
	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			log.Printf("Table %v does not exist.\n", tableName)
			err = nil
		} else {
			log.Printf("Couldn't determine existence of table %v. Here's why: %v\n", tableName, err)
		}
		exists = false
	}
	return exists, err
}

func (setup DynamoConfigurator) createMainTable() error {

	// create table
	tableInput := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("pk"),
				AttributeType: types.ScalarAttributeTypeS, // data type descriptor: S == string
			},
			{
				AttributeName: aws.String("sk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("pk"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("sk"),
				KeyType:       types.KeyTypeRange,
			},
		},
		TableName: aws.String(MainTableName),
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	}
	err := setup.createTable(MainTableName, tableInput)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

// createDynamoDBTable creates a table in the client's instance
// using the table parameters specified in input.
func (setup DynamoConfigurator) createTable(tableName string, input *dynamodb.CreateTableInput,
) error {
	var tableDesc *types.TableDescription
	table, err := setup.client.CreateTable(context.TODO(), input)
	if err != nil {
		log.Printf("Failed to create table %v with error: %v\n", tableName, err)
	} else {
		waiter := dynamodb.NewTableExistsWaiter(setup.client)
		err = waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{
			TableName: aws.String(tableName)}, 5*time.Minute)
		if err != nil {
			log.Printf("Failed to wait on create table %v with error: %v\n", tableName, err)
		}
		tableDesc = table.TableDescription
	}

	fmt.Println(tableDesc)

	return err
}
