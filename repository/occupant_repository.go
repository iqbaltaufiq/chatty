package repository

import "context"

type OccupantRepository interface {
	Add(ctx context.Context, roomID, username string) error
	Remove(ctx context.Context, roomID, username string) error
}
