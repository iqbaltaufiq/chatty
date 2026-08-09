package repository

import (
	"context"
	"database/sql"

	"chatty/entity"
)

type roomRepository struct{}

func NewRoomRepository() RoomRepository {
	return &roomRepository{}
}

// Create inserts a newly created room into db.
func (r *roomRepository) Create(ctx context.Context, dbtx DBTX, room entity.Room) error {
	query := `
	INSERT INTO rooms (id, name, is_private)
	VALUES (?,?,?)
	`
	_, err := dbtx.ExecContext(
		ctx,
		query,
		room.ID,
		room.Name,
		room.IsPrivate,
	)
	return err
}

// UpdateName alters the room's name.
func (r *roomRepository) UpdateName(ctx context.Context, dbtx DBTX, roomID, roomName string) error {
	query := `
	UPDATE rooms
	SET name = ?
	WHERE id = ?
	LIMIT 1
	`
	_, err := dbtx.ExecContext(ctx, query, roomName, roomID)
	return err
}

func (r *roomRepository) Exists(ctx context.Context, dbtx DBTX, roomID string) (bool, error) {
	query := `
	SELECT COUNT(*) as 'rows'
	FROM rooms
	WHERE id = ?
	LIMIT 1
	`
	rows := 0
	exists := false
	if err := dbtx.QueryRowContext(ctx, query, roomID).Scan(&rows); err != nil {
		return exists, err
	}
	if rows > 0 {
		exists = true
	}
	return exists, nil
}

func (r *roomRepository) Delete(ctx context.Context, dbtx DBTX, roomID string) error {
	query := `
	DELETE FROM rooms
	WHERE id = ?
	LIMIT 1
	`
	result, err := dbtx.ExecContext(ctx, query, roomID)
	if err != nil {
		return err
	}
	numRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if numRows == 0 {
		return sql.ErrNoRows
	}
	return err
}
