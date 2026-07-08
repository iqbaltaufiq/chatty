package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
	"time"

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
	allowedExt := imageExtensions
	photo = strings.TrimSpace(photo)
	splitFilename := strings.Split(photo, ".")
	length := len(splitFilename)

	// if there are more than one period (.) in the filename,
	// throw an error.
	if length != 2 {
		return errors.New("invalid filename")
	}

	nameWithoutExt := splitFilename[0]
	fileExt := splitFilename[1]

	if !allowedExt.Exists(fileExt) {
		return errors.New("extension not allowed")
	}

	// randomize the filename to avoid name collision
	// by hashing them with SHA256.
	// the composite: filename without extensions + username + nanoseconds
	nanoseconds := strconv.Itoa(time.Now().UTC().Nanosecond())
	stringToHash := nameWithoutExt + username + nanoseconds
	hashBytes := sha256.Sum256([]byte(stringToHash))
	hashString := hex.EncodeToString(hashBytes[:])
	finalFilename := hashString + "." + fileExt

	return s.userRepo.UpdatePhoto(ctx, username, finalFilename)
}
