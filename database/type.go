package database

import "time"

const (
	MAX_IDLE_CONNS int           = 10
	MAX_OPEN_CONNS int           = 50
	MAX_IDLE_TIME  time.Duration = 10 * time.Minute
	MAX_LIFETIME   time.Duration = 30 * time.Minute
)

type DatabaseConfig struct {
	Driver,
	Name,
	Username,
	Password,
	Host,
	Port string
}
