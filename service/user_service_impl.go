package service

import (
	"context"
	"database/sql"
	"io"
	"strings"

	"chatty/entity"
	"chatty/repository"
	"chatty/util"
)

type userService struct {
	db       *sql.DB
	userRepo repository.UserRepository
	fileRepo repository.FileRepository
}

func NewUserService(
	db *sql.DB,
	userRepo repository.UserRepository,
	fileRepo repository.FileRepository,
) UserService {
	return &userService{
		db,
		userRepo,
		fileRepo,
	}
}

// GetUser retrieves user based on their username.
func (s *userService) GetUser(ctx context.Context, username string) (entity.User, error) {
	return s.userRepo.GetUser(ctx, s.db, username)
}

// GetUserByUUID retrieves user based on their UUID.
// Note that their uuid is stored in BINARY(16), so a conversion from string to binary is necessary.
func (s *userService) GetUserByUUID(ctx context.Context, uuid string) (entity.User, error) {
	binaryUUID, err := util.UUIDtoBinary(uuid)
	if err != nil {
		return entity.User{}, err
	}
	return s.userRepo.GetUserByUUID(ctx, s.db, binaryUUID)
}

// UpdateFullname changes a user's full name.
func (s *userService) UpdateFullname(ctx context.Context, username, fullname string) error {
	fullname = strings.TrimSpace(fullname)
	return s.userRepo.UpdateFullname(ctx, s.db, username, fullname)
}

// UpdatePhoto changes a user's profile photo.
// `photo` is a filename complete with its file extension.
//
// valid: `default.png`, `16228591.jpg`
func (s *userService) UpdatePhoto(ctx context.Context, file io.Reader, username, photo string) error {
	hashedFilename, err := util.HashFilename(photo)
	if err != nil {
		return err
	}
	if err := s.userRepo.UpdatePhoto(ctx, s.db, username, hashedFilename); err != nil {
		return err
	}
	if err := s.fileRepo.Store(ctx, file, hashedFilename); err != nil {
		return err
	}
	return nil
}
