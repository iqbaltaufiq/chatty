package router

import (
	"database/sql"
	"net/http"

	mw "chatty/middleware"
)

func NewRouter(db *sql.DB) http.Handler {
	var (
		mux          = http.NewServeMux()
		repositories = NewRepositories(db)
		services     = NewServices(repositories)
		handlers     = NewHandlers(services)
		authMidw     = mw.NewAuthMiddleware(services.User())
	)

	authHandler := handlers.Auth()

	mux.Handle("POST /api/register", mw.Attach(authHandler.Register, authMidw.NoAuth))
	mux.Handle("POST /api/login", mw.Attach(authHandler.Login, authMidw.NoAuth))
	mux.Handle("POST /api/logout", mw.Attach(authHandler.Logout, authMidw.Protected))
	return mux
}
