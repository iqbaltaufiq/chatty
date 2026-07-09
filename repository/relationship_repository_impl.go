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

func (r *relationshipRepository) Remove(ctx context.Context, self, target string) error {
	query := `
	DELETE FROM relationships
	WHERE initiator = ? AND target = ? OR initiator = ? AND target = ?
	LIMIT 1
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	_, err := stmt.ExecContext(
		ctx,
		self, target,
		target, self,
	)
	return err
}

// Remove removes relationship between two users.
func (r *relationshipRepository) RemoveByID(ctx context.Context, relationshipID int) error {
	query := `
	DELETE FROM relationships
	WHERE id = ?
	LIMIT 1
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	_, err := stmt.ExecContext(ctx, relationshipID)
	return err
}

// Exists checks if two users have befriended with each other.
func (r *relationshipRepository) Exists(ctx context.Context, initiator, target string) (bool, error) {
	query := `
	SELECT COUNT(*) as 'rowsCount'
	FROM relationships
	WHERE initiator = ? AND target = ? OR initiator = ? AND target = ?
	LIMIT 1
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		initiator, target,
		target, initiator,
	)

	exists := false
	rowsCount := 0
	if err := row.Scan(&rowsCount); err != nil {
		return exists, err
	}
	if rowsCount > 0 {
		exists = true
	}
	return exists, nil
}

func (r *relationshipRepository) GetFriends(ctx context.Context, username string) ([]model.Friend, error) {
	// not yet implemented
	return nil, nil
}
