package service

import (
	"context"
	"database/sql"
	"strings"

	"chatty/model"
	"chatty/repository"
)

type relationshipService struct {
	db               *sql.DB
	relationshipRepo repository.RelationshipRepository
}

func NewRelationshipService(
	db *sql.DB,
	relationshipRepo repository.RelationshipRepository,
) RelationshipService {
	return &relationshipService{
		db,
		relationshipRepo,
	}
}

// Add befriends initiator with the target.
//
// Returns error if they are already friends.
func (s *relationshipService) Add(ctx context.Context, initiator, target string) error {
	initiator = strings.TrimSpace(initiator)
	target = strings.TrimSpace(target)

	tx, _ := s.db.Begin()
	defer tx.Rollback()

	exists, err := s.relationshipRepo.Exists(ctx, tx, initiator, target)
	if err != nil {
		return err
	}
	if exists {
		return &RowsErrAlreadyExists{}
	}
	if err := s.relationshipRepo.Add(ctx, tx, initiator, target); err != nil {
		return err
	}
	return tx.Commit()
}

// Remove unfriends both users.
//
// Returns error if they haven't become friends yet.
func (s *relationshipService) Remove(ctx context.Context, initiator, target string) error {
	initiator = strings.TrimSpace(initiator)
	target = strings.TrimSpace(target)

	tx, _ := s.db.Begin()
	defer tx.Rollback()

	exists, err := s.relationshipRepo.Exists(ctx, tx, initiator, target)
	if err != nil {
		return err
	}
	if !exists {
		return sql.ErrNoRows
	}
	if err := s.relationshipRepo.Remove(ctx, tx, initiator, target); err != nil {
		return err
	}
	return tx.Commit()
}

// RemoveByID unfriends two users referenced by their relationshipID.
func (s *relationshipService) RemoveByID(ctx context.Context, relationshipID int) error {
	return s.relationshipRepo.RemoveByID(ctx, nil, relationshipID)
}

// GetFriends retrieves a list of friends that belongs to the user.
func (s *relationshipService) GetFriends(ctx context.Context, username string) ([]model.Friend, error) {
	username = strings.TrimSpace(username)
	return s.relationshipRepo.GetFriends(ctx, nil, username)
}
