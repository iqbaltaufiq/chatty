package repository

import (
	"context"

	"chatty/model"
)

type RelationshipRepository interface {
	Add(ctx context.Context, initiator, target string) error
	Remove(ctx context.Context, self, target string) error
	RemoveByID(ctx context.Context, relationshipID int) error
	Exists(ctx context.Context, initiator, target string) (bool, error)
	GetFriends(ctx context.Context, username string) ([]model.Friend, error)
}
