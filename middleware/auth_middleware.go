package middleware

import (
	"context"
	"net/http"

	"chatty/handler"
	"chatty/service"
	"chatty/util"
)

type authMiddleware struct {
	userService service.UserService
}

func NewAuthMiddleware(userService service.UserService) *authMiddleware {
	return &authMiddleware{userService}
}

// Protected rejects unauthenticated request from accessing protected endpoint.
//
// It parses the JWT token from authentication cookie
// and eventually attaches user's uuid and username into the request's context.
func (m *authMiddleware) Protected(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie(util.AuthCookieName)
		if err != nil {
			handler.JSONError(w, http.StatusUnauthorized, err)
			return
		}

		// get UserID from jwt claims
		// and get the user's username from db.
		// ** note: username shouldn't be in jwt since it's a modifiable identifier. **
		claims, err := util.ValidateJWT(authCookie.Value)
		if err != nil {
			handler.JSONError(w, http.StatusUnauthorized, err)
			return
		}
		user, err := m.userService.GetUserByUUID(r.Context(), claims.UserID)
		if err != nil {
			handler.JSONError(w, http.StatusUnauthorized, err)
			return
		}
		username := user.Username

		// attach uuid and username to the context
		ctxWithUUID := context.WithValue(r.Context(), uuidKey, claims.UserID)
		finalCtx := context.WithValue(ctxWithUUID, usernameKey, username)

		next.ServeHTTP(w, r.WithContext(finalCtx))
	})
}
