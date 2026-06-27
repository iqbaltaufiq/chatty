package service

import (
	"context"

	"chatty/entity"
	"chatty/model"
)

type AuthService interface {
	CreateUser(ctx context.Context, user entity.User) error
	ValidateCreds(ctx context.Context, creds model.Credential) error
}
