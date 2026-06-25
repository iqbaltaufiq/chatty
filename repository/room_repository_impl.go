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

// Create inserts a newly created room into db.
func (r *roomRepository) Create(ctx context.Context, room entity.Room) error {
	query := `
	INSERT INTO rooms (id, name, is_private)
	VALUES (?,?,?)
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	_, err := stmt.ExecContext(
		ctx,
		room.ID,
		room.Name,
		room.IsPrivate,
	)
	return err
}

func (r *roomRepository) UpdateName(ctx context.Context, roomID, roomName string) error {
	// not yet implemented
	return nil
}
