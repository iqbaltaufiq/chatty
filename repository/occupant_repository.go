package repository

import (
	"context"
)

type OccupantRepository interface {
	GetRoomBetweenUsers(ctx context.Context, dbtx DBTX, firstUser, secondUser string) (string, error)
	Add(ctx context.Context, dbtx DBTX, roomID, username string) error
	Remove(ctx context.Context, dbtx DBTX, roomID, username string) error
	Exists(ctx context.Context, dbtx DBTX, roomID, username string) (bool, error)
}
