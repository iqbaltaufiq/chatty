package service

import (
	"context"

	"chatty/model"
	"chatty/repository"
)

type relationshipService struct {
	relationshipRepo repository.RelationshipRepository
}

func NewRelationshipService(relationshipRepo repository.RelationshipRepository) RelationshipService {
	return &relationshipService{relationshipRepo}
}

func (s *relationshipService) Add(ctx context.Context, initiator, target string) error {
	// not yet implemented
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
