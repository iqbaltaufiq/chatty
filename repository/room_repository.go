package repository

import (
	"context"

	"chatty/entity"
)

type RoomRepository interface {
	Create(ctx context.Context, dbtx DBTX, room entity.Room) error
	UpdateName(ctx context.Context, dbtx DBTX, roomID, roomName string) error
	Exists(ctx context.Context, dbtx DBTX, roomID string) (bool, error)
	Delete(ctx context.Context, dbtx DBTX, roomID string) error
}
