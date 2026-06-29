package router

import (
	"database/sql"
	"net/http"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	var (
		mux          = http.NewServeMux()
		repositories = NewRepositories(db)
		services     = NewServices(repositories)
		handlers     = NewHandlers(services)
	)

	mux.HandleFunc("POST /api/register", handlers.Auth().Register)
	mux.HandleFunc("POST /api/login", handlers.Auth().Login)

	return mux
}
