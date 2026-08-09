package repository

import (
	"context"

	"chatty/entity"
)

type UserRepository interface {
	GetUser(ctx context.Context, dbtx DBTX, username string) (entity.User, error)
	GetUserByUUID(ctx context.Context, dbtx DBTX, uuid []byte) (entity.User, error)
	UpdateFullname(ctx context.Context, dbtx DBTX, username, fullname string) error
	UpdatePhoto(ctx context.Context, dbtx DBTX, username, photo string) error

	// methods below are used in authentication service
	CreateUser(ctx context.Context, dbtx DBTX, user entity.User) error
	GetPassword(ctx context.Context, dbtx DBTX, username string) (string, error)
}
