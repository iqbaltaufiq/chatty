package middleware

import (
	"errors"
	"net/http"

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
			util.JSONError(w, http.StatusUnauthorized, err)
			return
		}

		// get UserID from jwt claims
		// and get the user's username from db.
		// ** note: username shouldn't be in jwt since it's a modifiable identifier. **
		claims, err := util.ValidateJWT(authCookie.Value)
		if err != nil {
			util.JSONError(w, http.StatusUnauthorized, err)
			return
		}
		user, err := m.userService.GetUserByUUID(r.Context(), claims.UserID)
		if err != nil {
			util.JSONError(w, http.StatusUnauthorized, err)
			return
		}
		username := user.Username

		// attach uuid and username to the context
		ctxWithUUID := util.SetCtxUUID(r.Context(), claims.UserID)
		finalCtx := util.SetCtxUsername(ctxWithUUID, username)

		next.ServeHTTP(w, r.Clone(finalCtx))
	})
}

// NoAuth rejects incoming request from authenticated user.
// It is different than a plain endpoint without middleware.
//
// This is the oppsite of `Protected` method, preventing endpoints like `/login` and `/register`
// being accessed by authenticated user.
func (m *authMiddleware) NoAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := r.Cookie(util.AuthCookieName); errors.Is(err, http.ErrNoCookie) {
			next.ServeHTTP(w, r)
			return
		}
		util.JSONError(w, http.StatusForbidden, &AuthErrAlreadyLoggedIn{})
	})
}
