package main

import (
	"chatty/database"
	"chatty/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	var (
		dbDriver   = os.Getenv("DB_DRIVER")
		dbName     = os.Getenv("DB_NAME")
		dbUsername = os.Getenv("DB_USERNAME")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		appHost    = os.Getenv("APP_HOST")
		appPort    = os.Getenv("APP_PORT")
		serverAddr = appHost + ":" + appPort
	)

	dbconfig := &database.DatabaseConfig{
		Driver:   dbDriver,
		Name:     dbName,
		Username: dbUsername,
		Password: dbPassword,
		Host:     dbHost,
		Port:     dbPort,
	}

	db, err := database.Connect(dbconfig)
	if err != nil {
		log.Fatalln("error initializing connection to database")
	}
	defer db.Close()

	router := router.NewRouter(db)
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
