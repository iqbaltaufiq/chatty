package util_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"chatty/util"
)

func TestJwt(test *testing.T) {
	userID := util.GenerateUUID()
	stringToken, err := util.GenerateJWT(userID, "user")
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		test.Fail()
	}

	test.Run("should be valid", func(t *testing.T) {
		host := os.Getenv("APP_HOST")
		port := os.Getenv("APP_PORT")
		issuer := fmt.Sprintf("http://%s:%s", host, port)

		claims, err := util.ValidateJWT(stringToken)
		if err != nil {
			log.Printf("failed when validating token: %s\n", err.Error())
			t.Fail()
		}
		if claims.Issuer != issuer {
			log.Printf("expected: %s, actual: %s\n", issuer, claims.Issuer)
			t.Fail()
		}
		if claims.UserID != userID {
			log.Println("userID does not exist ")
			t.Fail()
		}
		if time.Now().UTC().After(claims.ExpiresAt.Time) {
			log.Printf("token is expired: %v\n", claims.ExpiresAt.Time)
			t.Fail()
		}
	})
	log.Printf("token :%s\n", stringToken)
}
