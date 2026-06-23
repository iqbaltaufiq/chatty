package repository

import (
	"context"

	"chatty/entity"
)

type UserRepository interface {
	GetUser(ctx context.Context, username string) (entity.User, error)
	GetUserByUUID(ctx context.Context, uuid []byte) (entity.User, error)
	UpdateFullname(ctx context.Context, username, fullname string) error
	UpdatePhoto(ctx context.Context, username, photo string) error

	// methods below are used in authentication service
	CreateUser(ctx context.Context, user entity.User) error
	GetPassword(ctx context.Context, username string) (string, error)
}
