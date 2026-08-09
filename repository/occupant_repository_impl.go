package repository

import (
	"context"
)

type occupantRepository struct{}

func NewOccupantRepository() OccupantRepository {
	return &occupantRepository{}
}

// Add assigns user to the room.
func (r *occupantRepository) Add(ctx context.Context, dbtx DBTX, roomID, username string) error {
	query := `
	INSERT INTO occupants (room_id, username)
	VALUES (?,?)
	`
	_, err := dbtx.ExecContext(ctx, query, roomID, username)
	return err
}

// Remove takes out user from the room.
func (r *occupantRepository) Remove(ctx context.Context, dbtx DBTX, roomID, username string) error {
	query := `
	DELETE FROM occupants
	WHERE room_id = ? AND username = ?
	LIMIT 1
	`
	_, err := dbtx.ExecContext(ctx, query, roomID, username)
	return err
}

// Exists checks if user exists in the specified room
func (r *occupantRepository) Exists(ctx context.Context, dbtx DBTX, roomID, username string) (bool, error) {
	query := `
	SELECT COUNT(*) as 'rows'
	FROM occupants
	WHERE room_id = ? AND username = ?
	LIMIT 1
	`
	exists := false
	rowsFound := 0
	row := dbtx.QueryRowContext(ctx, query, roomID, username)
	if err := row.Scan(&rowsFound); err != nil {
		return exists, err
	}
	if rowsFound > 0 {
		exists = true
	}
	return exists, nil
}
