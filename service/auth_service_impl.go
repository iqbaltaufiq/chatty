package service

import (
	"context"
	"database/sql"
	"io"

	"chatty/entity"
	"chatty/model"
	customerrors "chatty/model/custom_errors"
	"chatty/repository"
	"chatty/util"
)

type authService struct {
	db       *sql.DB
	userRepo repository.UserRepository
	fileRepo repository.FileRepository
}

func NewAuthService(
	db *sql.DB,
	userRepo repository.UserRepository,
	fileRepo repository.FileRepository,
) AuthService {
	return &authService{
		db,
		userRepo,
		fileRepo,
	}
}

// CreateUser generates UUID, hashes the password, and stores the user into database.
func (s *authService) CreateUser(ctx context.Context, user entity.User, file io.Reader) error {
	var (
		usernameExists = user.Username != ""
		fullnameExists = user.Fullname != ""
		emailExists    = user.Email != ""
		passwordExists = user.Password != ""
		fileExists     = file != nil
	)
	if !usernameExists ||
		!fullnameExists ||
		!emailExists ||
		!passwordExists {
		return &customerrors.ErrInsufficientPayload{}
	}
	if fileExists {
		hashedFilename, err := util.HashFilename(user.Photo)
		if err != nil {
			return err
		}
		user.Photo = hashedFilename
	}
	// hash user's password using bcrypt
	hashPwd, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPwd
	user.UUID = util.GenerateUUID()

	if err := s.userRepo.CreateUser(ctx, s.db, user); err != nil {
		return err
	}
	if fileExists {
		if err := s.fileRepo.Store(
			ctx,
			file,
			user.Photo,
		); err != nil {
			return err
		}
	}
	return nil
}

// ValidateCreds checks if password from request matches the password stored in database.
func (s *authService) ValidateCreds(ctx context.Context, creds model.Credential) error {
	dbPassword, err := s.userRepo.GetPassword(ctx, s.db, creds.Username)
	if err != nil {
		return err
	}
	return util.MatchPassword(creds.Password, dbPassword)
}
