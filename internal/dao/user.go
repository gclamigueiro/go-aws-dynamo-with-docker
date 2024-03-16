package dao

import (
	"context"

	"github.com/gclamigueiro/go-aws-dynamo-with-docker/internal/models"
)

type UserDao interface {
	GetUser(ctx context.Context, username string) (models.User, error)
	AddUser(ctx context.Context, user models.User) error
}
