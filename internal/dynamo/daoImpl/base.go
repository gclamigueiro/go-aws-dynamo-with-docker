package daoImpl

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type baseDao struct {
	client *dynamodb.Client
}
