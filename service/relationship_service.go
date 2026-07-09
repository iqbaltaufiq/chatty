package service

import (
	"context"

	"chatty/model"
)

type RelationshipService interface {
	Add(ctx context.Context, initiator, target string) error
	Remove(ctx context.Context, initiator, target string) error
	RemoveByID(ctx context.Context, relationshipID int) error
	GetFriends(ctx context.Context, username string) ([]model.Friend, error)
}
