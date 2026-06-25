package repository

import (
	"context"
	"database/sql"

	"chatty/entity"
)

type roomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{db}
}

func (r *roomRepository) Create(ctx context.Context, room entity.Room) error {
	// not yet implemented
	return nil
}

func (r *roomRepository) UpdateName(ctx context.Context, roomID, roomName string) error {
	// not yet implemented
	return nil
}
