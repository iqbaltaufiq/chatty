package repository

import (
	"context"
	"database/sql"
)

type OccupantRepository interface {
	Add(ctx context.Context, tx *sql.Tx, roomID, username string) error
	Remove(ctx context.Context, tx *sql.Tx, roomID, username string) error
	Exists(ctx context.Context, tx *sql.Tx, roomID, username string) (bool, error)
}
