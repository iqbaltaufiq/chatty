package repository

import (
	"context"
	"database/sql"
)

type occupantRepository struct {
	db *sql.DB
}

func NewOccupantRepository(db *sql.DB) OccupantRepository {
	return &occupantRepository{db}
}

// Add assigns user to the room.
func (r *occupantRepository) Add(ctx context.Context, roomID, username string) error {
	// not yet implemented
	return nil
}

// Remove takes out user from the room.
func (r *occupantRepository) Remove(ctx context.Context, roomID, username string) error {
	// not yet implemented
	return nil
}
