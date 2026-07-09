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
		authHandler         = handlers.Auth()
		userHandler         = handlers.User()
		roomHandler         = handlers.Room()
		occupantHandler     = handlers.Occupant()
		relationshipHandler = handlers.Relationship()
	)

	mux.Handle("POST /api/register", mw.Attach(authHandler.Register, authMidw.NoAuth))
	mux.Handle("POST /api/login", mw.Attach(authHandler.Login, authMidw.NoAuth))
	mux.Handle("POST /api/logout", mw.Attach(authHandler.Logout, authMidw.Protected))

	mux.Handle("GET /api/user/{identifier}/get", mw.Attach(userHandler.Get, authMidw.Protected))
	mux.Handle("PATCH /api/user/fullname/update", mw.Attach(userHandler.UpdateFullname, authMidw.Protected))
	mux.Handle("PATCH /api/user/photo/update", mw.Attach(userHandler.UpdatePhoto, authMidw.Protected))

	mux.Handle("POST /api/room/new", mw.Attach(roomHandler.Create, authMidw.Protected))
	mux.Handle("PATCH /api/room/name/update", mw.Attach(roomHandler.UpdateName, authMidw.Protected))

	mux.Handle("POST /api/occupant/add", mw.Attach(occupantHandler.Add, authMidw.Protected))
	mux.Handle("DELETE /api/occupant/remove", mw.Attach(occupantHandler.Remove, authMidw.Protected))

	mux.Handle("GET /api/relationship/all", mw.Attach(relationshipHandler.GetFriends, authMidw.Protected))

	// attach global middleware to multiplexer
	muxWithGlobalMw := commonMw.AttachContext(mux)
	return muxWithGlobalMw
}
