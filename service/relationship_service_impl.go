package service

import (
	"context"
	"strings"

	"chatty/model"
	"chatty/repository"
)

type relationshipService struct {
	relationshipRepo repository.RelationshipRepository
}

func NewRelationshipService(relationshipRepo repository.RelationshipRepository) RelationshipService {
	return &relationshipService{relationshipRepo}
}

// Add befriends initiator with the target.
//
// Returns error if they are already friends.
func (s *relationshipService) Add(ctx context.Context, initiator, target string) error {
	initiator = strings.TrimSpace(initiator)
	target = strings.TrimSpace(target)

	exists, err := s.relationshipRepo.Exists(ctx, initiator, target)
	if err != nil {
		return err
	}
	if exists {
		return &RowsErrAlreadyExists{}
	}
	return s.relationshipRepo.Add(ctx, initiator, target)
}

func (s *relationshipService) Remove(ctx context.Context, initiator, target string) error {
	// not yet implemented
}

func (s *relationshipService) RemoveByID(ctx context.Context, relationshipID int) error {
	// not yet implemented
}

func (s *relationshipService) GetFriends(ctx context.Context, username string) ([]model.Friend, error) {
	// not yet implemented
}
