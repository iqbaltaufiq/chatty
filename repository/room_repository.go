package repository

import (
	"context"

	"chatty/entity"
)

type RoomRepository interface {
	Create(ctx context.Context, room entity.Room) error
	UpdateName(ctx context.Context, roomID, roomName string) error
}
