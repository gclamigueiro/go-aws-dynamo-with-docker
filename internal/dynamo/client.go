package dynamo

import (
	"context"
	"fmt"

	config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

type clientOptions func(*dynamodb.Options)

func GetLocalConfiguration(options *dynamodb.Options) {
	endpoint := "http://localhost:8000"
	options.Region = "us-west-2"
	options.Credentials = credentials.NewStaticCredentialsProvider("local", "local", "local")
	options.BaseEndpoint = aws.String(endpoint)
}

func NewDynamoDBClient(opts ...clientOptions) (*dynamodb.Client, error) {

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create aws config: %w", err)
	}

	client := dynamodb.NewFromConfig(cfg, func(options *dynamodb.Options) {
		// Apply supplied options
		for _, fn := range opts {
			fn(options)
		}
	})

	return client, nil
}
