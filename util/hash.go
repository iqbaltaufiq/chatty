package util

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var allowedExt = map[string]struct{}{
	"jpg":  {},
	"jpeg": {},
	"png":  {},
	"webp": {},
	"avif": {},
	"tiff": {},
}

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
	// if there are more than one period (.) in the filename,
	// throw an error.
	splitFilename := strings.Split(strings.TrimSpace(filename), ".")
	if len(splitFilename) != 2 {
		return "", errors.New("invalid filename")
	}
	// Check if the file extension is allowed.
	fileExt := splitFilename[1]
	if _, ok := allowedExt[fileExt]; !ok {
		return "", errors.New("extension not allowed")
	}
	// randomize the filename to avoid name collision
	// by hashing them with SHA1.
	nameWithoutExt := splitFilename[0]
	randAlphanumeric := GenerateAlphanumeric(8)

	stringComposite := nameWithoutExt + randAlphanumeric

	hashBytes := sha1.Sum([]byte(stringComposite))
	hashString := hex.EncodeToString(hashBytes[:])
	finalFilename := hashString + "." + fileExt

	return finalFilename, nil
}
