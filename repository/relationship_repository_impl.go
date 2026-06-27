package repository

import (
	"context"
	"database/sql"

	"chatty/model"
)

type relationshipRepository struct {
	db *sql.DB
}

func NewRelationshipRepository(db *sql.DB) RelationshipRepository {
	return &relationshipRepository{db}
}

// Add befriends initiator with the target.
func (r *relationshipRepository) Add(ctx context.Context, initiator, target string) error {
	query := `
	INSERT INTO relationships (initiator, target)
	VALUES (?,?)
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	_, err := stmt.ExecContext(ctx, initiator, target)
	return err
}

func (r *relationshipRepository) Remove(ctx context.Context, initiator, target string) error {
	// not yet implemented
	return nil
}

func (r *relationshipRepository) Exists(ctx context.Context, initiator, target string) (bool, error) {
	// not yet implemented
	return false, nil
}

func (r *relationshipRepository) GetFriends(ctx context.Context, username string) ([]model.Friend, error) {
	// not yet implemented
	return nil, nil
}
