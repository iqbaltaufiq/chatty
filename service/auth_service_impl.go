package service

import (
	"context"

	"chatty/entity"
	"chatty/model"
	"chatty/repository"
	"chatty/util"
)

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo,
	}
}

// CreateUser generates UUID, hashes the password, and stores the user into database.
func (s *authService) CreateUser(ctx context.Context, user entity.User) error {
	// hash user's password using bcrypt
	hashPwd, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPwd
	user.UUID = util.GenerateUUID()

	return s.userRepo.CreateUser(ctx, user)
}

// ValidateCreds checks if password from request matches the password stored in database.
func (s *authService) ValidateCreds(ctx context.Context, creds model.Credential) error {
	dbPassword, err := s.userRepo.GetPassword(ctx, creds.Username)
	if err != nil {
		return err
	}
	return util.MatchPassword(creds.Password, dbPassword)
}
