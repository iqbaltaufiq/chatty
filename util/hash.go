package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hashBytes), err
}

func MatchPassword(requestPwd, databasePwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(databasePwd), []byte(requestPwd))
}
