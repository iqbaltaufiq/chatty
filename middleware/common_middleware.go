package middleware

import (
	"net/http"

	"chatty/service"
	"chatty/util"
)

type commonMiddleware struct {
	userService service.UserService
}

func NewCommonMiddleware(userService service.UserService) *commonMiddleware {
	return &commonMiddleware{userService}
}

// AttachContext attaches userID and username to request's context.
//
// It does not reject the request even if the incoming request is unauthenticated.
// If you want to reject unauthenticated request, use AuthMiddleware.Protected() instead.
func (m *commonMiddleware) AttachContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie(util.AuthCookieName)
		if err != nil || authCookie.Valid() != nil {
			next.ServeHTTP(w, r)
			return
		}
		claims, err := util.ValidateJWT(authCookie.Value)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		user, err := m.userService.GetUserByUUID(r.Context(), claims.UserID)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		username := user.Username

		ctxWithUserID := SetCtxUUID(r.Context(), claims.UserID)
		ctxWithUsername := SetCtxUsername(ctxWithUserID, username)
		next.ServeHTTP(w, r.Clone(ctxWithUsername))
	})
}
