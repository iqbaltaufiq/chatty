package service

import (
	"context"
	"strings"

	"chatty/entity"
	"chatty/repository"
	"chatty/util"
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
	roomID := ""
	roomIDLength := 6
	for {
		roomID = util.GenerateAlphanumeric(roomIDLength)
		exists, _ := s.roomRepo.Exists(ctx, roomID)
		if !exists {
			break
		}
	}
	room.ID = roomID
	room.Name = strings.TrimSpace(room.Name)
	room.IsPrivate = true
	return roomID, s.roomRepo.Create(ctx, room)
}

// UpdateName changes the name of the room.
func (s *roomService) UpdateName(ctx context.Context, roomID, roomName string) error {
	roomName = strings.TrimSpace(roomName)
	roomID = strings.TrimSpace(roomID)

	if len(roomID) != 6 {
		return &RoomErrInvalidRoomID{}
	}
	return s.roomRepo.UpdateName(ctx, roomID, roomName)
}
