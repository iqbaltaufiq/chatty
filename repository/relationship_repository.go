package repository

import (
	"context"

	"chatty/entity"
	"chatty/model"
)

type RelationshipRepository interface {
	Add(ctx context.Context, dbtx DBTX, initiator, target string) error
	Remove(ctx context.Context, dbtx DBTX, self, target string) error
	RemoveByID(ctx context.Context, dbtx DBTX, relationshipID int) error
	Exists(ctx context.Context, dbtx DBTX, initiator, target string) (bool, error)
	Get(ctx context.Context, dbtx DBTX, id int) (entity.Relationship, error)
	GetFriends(ctx context.Context, dbtx DBTX, username string) ([]model.Friend, error)
}
