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

		// middlewares
		authMidw = mw.NewAuthMiddleware(services.User())
		commonMw = mw.NewCommonMiddleware(services.User())

		// handlers
		authHandler = handlers.Auth()
		userHandler = handlers.User()
	)

	mux.Handle("POST /api/register", mw.Attach(authHandler.Register, authMidw.NoAuth))
	mux.Handle("POST /api/login", mw.Attach(authHandler.Login, authMidw.NoAuth))
	mux.Handle("POST /api/logout", mw.Attach(authHandler.Logout, authMidw.Protected))

	mux.Handle("GET /api/user/{identifier}/get", mw.Attach(userHandler.Get, authMidw.Protected))
	mux.Handle("POST /api/user/fullname/update", mw.Attach(userHandler.UpdateFullname, authMidw.Protected))
	mux.Handle("PATCH /api/user/photo/update", mw.Attach(userHandler.UpdatePhoto, authMidw.Protected))

	// attach global middleware to multiplexer
	muxWithGlobalMw := commonMw.AttachContext(mux)
	return muxWithGlobalMw
}
