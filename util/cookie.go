package util

import (
	"net/http"
	"os"
)

const AuthCookieName = "auth_v1"

func SetAuthCookie(value string, maxAge int) *http.Cookie {
	appEnv := os.Getenv("APP_ENV")
	secureCookie := appEnv == "production"

	return &http.Cookie{
		Name:     AuthCookieName,
		Value:    value,
		MaxAge:   maxAge,
		Path:     "/",
		Secure:   secureCookie,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}
