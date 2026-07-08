package service

import "context"

type OccupantService interface {
	Add(ctx context.Context, roomID, username string) error
	Remove(ctx context.Context, roomID, username string) error
}
