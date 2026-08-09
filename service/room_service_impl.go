package service

import (
	"context"
	"database/sql"
	"strings"

	"chatty/entity"
	"chatty/repository"
	"chatty/util"
)

type roomService struct {
	db       *sql.DB
	roomRepo repository.RoomRepository
}

func NewRoomService(
	db *sql.DB,
	roomRepo repository.RoomRepository,
) RoomService {
	return &roomService{db, roomRepo}
}

// Create generates 6-digits unique ID for room ID
// and stores it in database.
// Unique ID must be uppercase alphanumeric.
//
// Returns the unique roomID and nil if success.
func (s *roomService) Create(ctx context.Context, room entity.Room) (string, error) {
	roomID := ""
	roomIDLength := 6
	tx, _ := s.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	for {
		roomID = util.GenerateAlphanumeric(roomIDLength)
		exists, _ := s.roomRepo.Exists(ctx, tx, roomID)
		if !exists {
			break
		}
	}
	room.ID = roomID
	if room.Name == "" {
		room.Name = "New group"
	} else {
		room.Name = strings.TrimSpace(room.Name)
	}
	room.IsPrivate = true
	if err := s.roomRepo.Create(ctx, tx, room); err != nil {
		return "", err
	}
	return roomID, tx.Commit()
}

// UpdateName changes the name of the room.
func (s *roomService) UpdateName(ctx context.Context, roomID, roomName string) error {
	roomName = strings.TrimSpace(roomName)
	roomID = strings.TrimSpace(roomID)

	if len(roomID) != 6 {
		return &RoomErrInvalidRoomID{}
	}
	tx, _ := s.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	exists, err := s.roomRepo.Exists(ctx, tx, roomID)
	if err != nil || !exists {
		return sql.ErrNoRows
	}
	if err := s.roomRepo.UpdateName(ctx, tx, roomID, roomName); err != nil {
		return err
	}
	return tx.Commit()
}
