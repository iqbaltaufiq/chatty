package service

import (
	"context"

	"chatty/entity"
	"chatty/repository"
)

type roomService struct {
	roomRepo repository.RoomRepository
}

func NewRoomService(roomRepo repository.RoomRepository) RoomService {
	return &roomService{roomRepo}
}

// Create generates 6-digits unique ID for room ID
// and stores it in database.
// Unique ID must be uppercase alphanumeric.
//
// Returns the unique roomID and nil if success.
func (s *roomService) Create(ctx context.Context, room entity.Room) (string, error) {
	// not yet implemented
}

// UpdateName changes the name of the room.
func (s *roomService) UpdateName(ctx context.Context, roomID, roomName string) error {
	// not yet implemented
}
