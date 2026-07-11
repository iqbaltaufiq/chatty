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
		authMw   = mw.NewAuthMiddleware(services.User())
		commonMw = mw.NewCommonMiddleware(services.User())

		// handlers
		authHandler         = handlers.Auth()
		userHandler         = handlers.User()
		roomHandler         = handlers.Room()
		occupantHandler     = handlers.Occupant()
		relationshipHandler = handlers.Relationship()
		chatHandler         = handlers.Chat()
	)

	mux.Handle("POST /api/register", mw.Attach(authHandler.Register, authMw.NoAuth))
	mux.Handle("POST /api/login", mw.Attach(authHandler.Login, authMw.NoAuth))
	mux.Handle("POST /api/logout", mw.Attach(authHandler.Logout, authMw.Protected))

	mux.Handle("GET /api/user/{identifier}/get", mw.Attach(userHandler.Get, authMw.Protected))
	mux.Handle("PATCH /api/user/fullname/update", mw.Attach(userHandler.UpdateFullname, authMw.Protected))
	mux.Handle("PATCH /api/user/photo/update", mw.Attach(userHandler.UpdatePhoto, authMw.Protected))

	mux.Handle("POST /api/room/new", mw.Attach(roomHandler.Create, authMw.Protected))
	mux.Handle("PATCH /api/room/name/update", mw.Attach(roomHandler.UpdateName, authMw.Protected))

	mux.Handle("POST /api/occupant/add", mw.Attach(occupantHandler.Add, authMw.Protected))
	mux.Handle("DELETE /api/occupant/remove", mw.Attach(occupantHandler.Remove, authMw.Protected))

	mux.Handle("GET /api/relationship/all", mw.Attach(relationshipHandler.GetFriends, authMw.Protected))
	mux.Handle("POST /api/relationship/{username}/add", mw.Attach(relationshipHandler.Befriend, authMw.Protected))
	mux.Handle("DELETE /api/relationship/{username}/remove", mw.Attach(relationshipHandler.Unfriend, authMw.Protected))
	mux.Handle("DELETE /api/relationship/id/{relationshipID}/remove", mw.Attach(relationshipHandler.UnfriendByID, authMw.Protected))

	mux.Handle("GET /api/chat/{roomID}/all", mw.Attach(chatHandler.GetAll, authMw.Protected))

	// attach global middleware to multiplexer
	muxWithGlobalMw := commonMw.AttachContext(mux)
	return muxWithGlobalMw
}
