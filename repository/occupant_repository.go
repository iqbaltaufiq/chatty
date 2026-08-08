package repository

import (
	"context"
	"database/sql"
)

type OccupantRepository interface {
	Add(ctx context.Context, tx *sql.Tx, roomID, username string) (*sql.Tx, error)
	Remove(ctx context.Context, tx *sql.Tx, roomID, username string) (*sql.Tx, error)
	Exists(ctx context.Context, tx *sql.Tx, roomID, username string) (*sql.Tx, bool, error)
}
