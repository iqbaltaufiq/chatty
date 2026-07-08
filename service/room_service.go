package service

import (
	"context"

	"chatty/entity"
)

type RoomService interface {
	Create(ctx context.Context, room entity.Room) error
	UpdateName(ctx context.Context, roomID, roomName string) error
}
