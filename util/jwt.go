package util

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	SigningMethod = jwt.SigningMethodHS256
	SecretKey     = os.Getenv("JWT_SECRET")
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT creates a jwt token used for authentication process.
// It will be stored as cookie value and intercepted by middleware for validation.
func GenerateJWT(userID, role string) (string, error) {
	var (
		host      = os.Getenv("APP_HOST")
		port      = os.Getenv("APP_PORT")
		issuer    = fmt.Sprintf("http://%s:%s", host, port)
		audience  = issuer
		issuedAt  = time.Now().UTC()
		expiresAt = issuedAt.Add(24 * time.Hour)
		jwtID     = uuid.Must(uuid.NewRandom()).String()
	)

	claims := CustomClaims{
		userID,
		role,
		jwt.RegisteredClaims{
			ID:        jwtID,
			Issuer:    issuer,
			Audience:  jwt.ClaimStrings{audience},
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	return jwt.NewWithClaims(SigningMethod, claims).
		SignedString([]byte(SecretKey))
}

// ValidateJWT checks if the passed token is a valid JWT token.
// Returns type-asserted `*CustomClaims` and nil if succeed.
func ValidateJWT(t string) (*CustomClaims, error) {
	secret := func(token *jwt.Token) (any, error) {
		return []byte(SecretKey), nil
	}

	token, err := jwt.ParseWithClaims(t, &CustomClaims{}, secret)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims, nil
}
