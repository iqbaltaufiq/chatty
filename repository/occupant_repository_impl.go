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
func (r *occupantRepository) Add(ctx context.Context, tx *sql.Tx, roomID, username string) (*sql.Tx, error) {
	query := `
	INSERT INTO occupants (room_id, username)
	VALUES (?,?)
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return tx, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, roomID, username)
	return tx, err
}

// Remove takes out user from the room.
func (r *occupantRepository) Remove(ctx context.Context, tx *sql.Tx, roomID, username string) (*sql.Tx, error) {
	query := `
	DELETE FROM occupants
	WHERE room_id = ? AND username = ?
	LIMIT 1
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return tx, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, roomID, username)
	return tx, err
}

// Exists checks if user exists in the specified room
func (r *occupantRepository) Exists(ctx context.Context, tx *sql.Tx, roomID, username string) (*sql.Tx, bool, error) {
	query := `
	SELECT COUNT(*) as 'rows'
	FROM occupants
	WHERE room_id = ? AND username = ?
	LIMIT 1
	`
	exists := false
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return tx, exists, err
	}
	defer stmt.Close()

	rowsFound := 0
	row := stmt.QueryRowContext(ctx, roomID, username)
	if err := row.Scan(&rowsFound); err != nil {
		return tx, exists, err
	}
	if rowsFound > 0 {
		exists = true
	}
	return tx, exists, nil
}
