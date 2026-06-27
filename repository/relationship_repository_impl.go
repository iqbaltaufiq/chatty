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

func (r *relationshipRepository) Add(ctx context.Context, initiator, target string) error {
	// not yet implemented
	return nil
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
