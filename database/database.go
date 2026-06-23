package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	driver := os.Getenv("DB_DRIVER")
	dsn := os.Getenv("DB_DSN")
	if driver == "" || dsn == "" {
		return nil, &ErrNoConfig{}
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(MAX_OPEN_CONNS)
	db.SetMaxIdleConns(MAX_IDLE_CONNS)
	db.SetConnMaxIdleTime(MAX_IDLE_TIME)
	db.SetConnMaxLifetime(MAX_LIFETIME)

	return db, nil
}
