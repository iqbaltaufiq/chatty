package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(config DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := sql.Open(config.Driver, dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(MAX_OPEN_CONNS)
	db.SetMaxIdleConns(MAX_IDLE_CONNS)
	db.SetConnMaxIdleTime(MAX_IDLE_TIME)
	db.SetConnMaxLifetime(MAX_LIFETIME)

	return db, nil
}
