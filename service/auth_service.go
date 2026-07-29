package service

import (
	"context"
	"io"

	"chatty/entity"
	"chatty/model"
)

type AuthService interface {
	CreateUser(ctx context.Context, user entity.User, file io.Reader) error
	ValidateCreds(ctx context.Context, creds model.Credential) error
}
