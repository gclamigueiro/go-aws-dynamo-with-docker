package models

type User struct {
	Username string `dynamodbav:"username"`
	Name     string `dynamodbav:"name"`
	LastName string `dynamodbav:"lastname"`
}
