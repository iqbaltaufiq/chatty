package util

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hashBytes), err
}

func MatchPassword(requestPwd, databasePwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(databasePwd), []byte(requestPwd))
}

// HashFilename is a filename complete with its file extension.
//
// valid: default.png, 16228591.jpg
func HashFilename(filename string) (string, error) {
	allowedExt := map[string]struct{}{
		"jpg":  {},
		"jpeg": {},
		"png":  {},
		"webp": {},
		"avif": {},
		"tiff": {},
	}
	// if there are more than one period (.) in the filename,
	// throw an error.
	splitFilename := strings.Split(strings.TrimSpace(filename), ".")
	if len(splitFilename) != 2 {
		return "", errors.New("invalid filename")
	}
	// Check if the file extension is allowed.
	nameWithoutExt := splitFilename[0]
	fileExt := splitFilename[1]
	if _, ok := allowedExt[fileExt]; !ok {
		return "", errors.New("extension not allowed")
	}
	// randomize the filename to avoid name collision
	// by hashing them with SHA256.
	// the composite: filename without extensions + random 3 alphanumeric + nanoseconds
	nanoseconds := strconv.Itoa(time.Now().UTC().Nanosecond())
	randAlphanumeric := GenerateAlphanumeric(3)
	stringComposite := nameWithoutExt + randAlphanumeric + nanoseconds

	hashBytes := sha256.Sum256([]byte(stringComposite))
	hashString := hex.EncodeToString(hashBytes[:])
	finalFilename := hashString + "." + fileExt

	return finalFilename, nil
}
