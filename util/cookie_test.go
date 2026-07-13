package util_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"chatty/util"
)

func TestAuthCookie(test *testing.T) {
	oneDay := 60 * 60 * 24
	authCookie := util.SetAuthCookie("jwt_token", oneDay)

	test.Run("should not expire", func(t *testing.T) {
		if authCookie.MaxAge == 0 {
			log.Println("cookie does not have expiration time")
			t.Fail()
		}
	})

	test.Run("should have correct path", func(t *testing.T) {
		if authCookie.Path != "/" {
			log.Println("path is not set correctly")
			t.Fail()
		}
	})

	test.Run("should be secure in production environment", func(t *testing.T) {
		appEnv := os.Getenv("APP_ENV")

		if appEnv == "production" && !authCookie.Secure {
			log.Println("cookie is not secure in production environment")
			t.Fail()
		}
	})

	test.Run("should not be accessible by client", func(t *testing.T) {
		if authCookie.HttpOnly != true {
			log.Println("cookie is not HttpOnly and can be modified by the client")
			t.Fail()
		}
	})

	test.Run("should only be available on the same domain", func(t *testing.T) {
		if authCookie.SameSite != http.SameSiteStrictMode {
			log.Println("cookie should be set to strict mode")
			t.Fail()
		}
	})

	test.Run("should set the maxage to negative on deletion", func(t *testing.T) {
		authCookie = util.DeleteCookie(util.AuthCookieName)

		if authCookie.MaxAge >= 0 {
			log.Println("max age is not set to negative")
			log.Printf("maxage: %v\n", authCookie.MaxAge)
			t.Fail()
		}
	})
}
