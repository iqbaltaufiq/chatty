package service

import (
	"context"
	"io"

	"chatty/entity"
	"chatty/model"
	"chatty/repository"
	"chatty/util"
)

type authService struct {
	userRepo repository.UserRepository
	fileRepo repository.FileRepository
}

func NewAuthService(
	userRepo repository.UserRepository,
	fileRepo repository.FileRepository,
) AuthService {
	return &authService{
		userRepo,
		fileRepo,
	}
}

// CreateUser generates UUID, hashes the password, and stores the user into database.
func (s *authService) CreateUser(ctx context.Context, user entity.User, file io.Reader) error {
	hashedFilename, err := util.HashFilename(user.Photo)
	if err != nil {
		return err
	}
	user.Photo = hashedFilename

	if err := s.fileRepo.Store(
		ctx,
		file,
		hashedFilename,
	); err != nil {
		return err
	}
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
