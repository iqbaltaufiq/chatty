package service

import (
	"context"
	"strings"

	"chatty/entity"
	"chatty/repository"
	"chatty/util"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

// GetUser retrieves user based on their username.
func (s *userService) GetUser(ctx context.Context, username string) (entity.User, error) {
	return s.userRepo.GetUser(ctx, username)
}

// GetUserByUUID retrieves user based on their UUID.
// Note that their uuid is stored in BINARY(16), so a conversion from string to binary is necessary.
func (s *userService) GetUserByUUID(ctx context.Context, uuid string) (entity.User, error) {
	binaryUUID, err := util.UUIDtoBinary(uuid)
	if err != nil {
		return entity.User{}, err
	}
	return s.userRepo.GetUserByUUID(ctx, binaryUUID)
}

// UpdateFullname changes a user's full name.
func (s *userService) UpdateFullname(ctx context.Context, username, fullname string) error {
	fullname = strings.TrimSpace(fullname)
	return s.userRepo.UpdateFullname(ctx, username, fullname)
}

// UpdatePhoto changes a user's profile photo.
// `photo` is a filename complete with its file extension.
//
// valid: `default.png`, `16228591.jpg`
func (s *userService) UpdatePhoto(ctx context.Context, username, photo string) error {
	hashedFilename, err := util.HashFilename(photo)
	if err != nil {
		return err
	}
	return s.userRepo.UpdatePhoto(ctx, username, hashedFilename)
}
