package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"chatty/database"
	"chatty/handler/bidirectional"
	"chatty/router"
)

func init() {
	godotenv.Load()
}

func main() {
	var (
		appHost    = os.Getenv("APP_HOST")
		appPort    = os.Getenv("APP_PORT")
		serverAddr = appHost + ":" + appPort
	)

	db, err := database.Connect()
	if err != nil {
		log.Fatalln("error initializing connection to database")
	}
	defer db.Close()

	websocketHub := bidirectional.NewHub()

	router := router.NewRouter(db, websocketHub)
	server := new(http.Server)
	server.Addr = serverAddr
	server.Handler = router

	log.Printf("starting the server on port %s...\n", appPort)
	if err := server.ListenAndServe(); err != nil {
		// Fatal will not run deferred function
		db.Close()
		log.Fatalf("server has stopped running. Reason: %s\n", err.Error())
	}
}
