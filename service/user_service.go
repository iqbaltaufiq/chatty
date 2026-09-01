package service

import (
	"context"

	"chatty/entity"
)

type UserService interface {
	GetUser(ctx context.Context, username string) (entity.User, error)
	GetUserByUUID(ctx context.Context, uuid string) (entity.User, error)
	UpdateFullname(ctx context.Context, username, fullname string) error
	UpdatePhoto(ctx context.Context, username, filename string) error
}
