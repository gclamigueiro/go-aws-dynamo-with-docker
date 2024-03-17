package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gclamigueiro/go-aws-dynamo-with-docker/internal/dynamo"
	"github.com/gclamigueiro/go-aws-dynamo-with-docker/internal/dynamo/daoImpl"
	"github.com/gclamigueiro/go-aws-dynamo-with-docker/internal/models"
)

func main() {

	endpoint := os.Getenv("DYNAMO_ENDPOINT")

	if endpoint == "" {
		endpoint = "http://localhost:8000"
	}

	client, err := dynamo.NewDynamoDBClient(dynamo.GetLocalConfiguration(endpoint))
	dynamoConfigurator := dynamo.NewDynamoConfigurator(client)

	if err != nil {
		panic(err)
	}

	// Setting up tables
	dynamoConfigurator.SetupDatabase()

	// Create some users
	userDao := daoImpl.NewUserDao(client)

	ctx := context.Background()

	userDao.AddUser(ctx, models.User{
		Username: "firstuser",
		Name:     "First",
		LastName: "User",
	})

	userDao.AddUser(ctx, models.User{
		Username: "seconduser",
		Name:     "Second",
		LastName: "User",
	})

	userDao.AddUser(ctx, models.User{
		Username: "thirduser",
		Name:     "Third",
		LastName: "User",
	})

	// Get User
	u, err := userDao.GetUser(ctx, "seconduser")
	fmt.Println(u, err)
}
