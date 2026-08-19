package service

import (
	"context"
	"database/sql"
	"strings"

	"chatty/entity"
	"chatty/model"
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
func (s *userService) UpdatePhoto(ctx context.Context, photo model.UploadedPhoto) error {
	if err := photo.IsValid(); err != nil {
		return err
	}
	hashedFilename, err := util.HashFilename(photo.GetFilename())
	if err != nil {
		return err
	}
	if err := s.userRepo.UpdatePhoto(ctx, s.db, photo.Username, hashedFilename); err != nil {
		return err
	}
	if err := s.fileRepo.Store(ctx, photo.File, hashedFilename); err != nil {
		return err
	}
	if photo.OldPhoto != "default.png" {
		if err := s.fileRepo.Delete(ctx, photo.OldPhoto); err != nil {
			return err
		}
	}
	return nil
}
