package repository

import (
	"context"
	"database/sql"
)

type occupantRepository struct{}

func NewOccupantRepository() OccupantRepository {
	return &occupantRepository{}
}

// Add assigns user to the room.
func (r *occupantRepository) Add(ctx context.Context, tx *sql.Tx, roomID, username string) error {
	query := `
	INSERT INTO occupants (room_id, username)
	VALUES (?,?)
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, roomID, username)
	return err
}

// Remove takes out user from the room.
func (r *occupantRepository) Remove(ctx context.Context, tx *sql.Tx, roomID, username string) error {
	query := `
	DELETE FROM occupants
	WHERE room_id = ? AND username = ?
	LIMIT 1
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, roomID, username)
	return err
}

// Exists checks if user exists in the specified room
func (r *occupantRepository) Exists(ctx context.Context, tx *sql.Tx, roomID, username string) (bool, error) {
	query := `
	SELECT COUNT(*) as 'rows'
	FROM occupants
	WHERE room_id = ? AND username = ?
	LIMIT 1
	`
	exists := false
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return exists, err
	}
	defer stmt.Close()

	rowsFound := 0
	row := stmt.QueryRowContext(ctx, roomID, username)
	if err := row.Scan(&rowsFound); err != nil {
		return exists, err
	}
	if rowsFound > 0 {
		exists = true
	}
	return exists, nil
}
