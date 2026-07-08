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

// UpdateName alters the room's name.
func (r *roomRepository) UpdateName(ctx context.Context, roomID, roomName string) error {
	query := `
	UPDATE rooms
	SET name = ?
	WHERE id = ?
	LIMIT 1
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	_, err := stmt.ExecContext(ctx, roomName, roomID)
	return err
}

func (r *roomRepository) Exists(ctx context.Context, roomID string) (bool, error) {
	query := `
	SELECT COUNT(*) as 'rows'
	FROM rooms
	WHERE id = ?
	LIMIT 1
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	rows := 0
	exists := false
	if err := stmt.QueryRowContext(ctx, roomID).Scan(&rows); err != nil {
		return exists, err
	}
	if rows > 0 {
		exists = true
	}
	return exists, nil
}
