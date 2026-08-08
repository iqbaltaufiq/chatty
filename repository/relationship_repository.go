package repository

import (
	"context"
	"database/sql"

	"chatty/model"
)

type RelationshipRepository interface {
	Add(ctx context.Context, tx *sql.Tx, initiator, target string) error
	Remove(ctx context.Context, tx *sql.Tx, self, target string) error
	RemoveByID(ctx context.Context, tx *sql.Tx, relationshipID int) error
	Exists(ctx context.Context, tx *sql.Tx, initiator, target string) (bool, error)
	GetFriends(ctx context.Context, tx *sql.Tx, username string) ([]model.Friend, error)
}
