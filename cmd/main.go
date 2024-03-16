package main

import (
	"context"
	"fmt"

	"github.com/gclamigueiro/go-aws-dynamo-with-docker/internal/dynamo"
	"github.com/gclamigueiro/go-aws-dynamo-with-docker/internal/dynamo/daoImpl"
	"github.com/gclamigueiro/go-aws-dynamo-with-docker/internal/models"
)

func main() {

	ctx := context.Background()
	client, err := dynamo.NewDynamoDBClient(dynamo.GetLocalConfiguration)
	dynamoConfigurator := dynamo.NewDynamoConfigurator(client)

	if err != nil {
		panic(err)
	}

	// Setting up tables
	dynamoConfigurator.SetupDatabase()

	// Create some users
	userDao := daoImpl.NewUserDao(client)

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
