package repository

import (
	"context"
)

type OccupantRepository interface {
	Add(ctx context.Context, dbtx DBTX, roomID, username string) error
	Remove(ctx context.Context, dbtx DBTX, roomID, username string) error
	Exists(ctx context.Context, dbtx DBTX, roomID, username string) (bool, error)
}
