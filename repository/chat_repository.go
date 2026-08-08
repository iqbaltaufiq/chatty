package repository

import (
	"context"
	"database/sql"

	"chatty/entity"
	"chatty/model"
)

type ChatRepository interface {
	Save(ctx context.Context, tx *sql.Tx, chat entity.Chat) (entity.Chat, error)
	GetAll(ctx context.Context, tx *sql.Tx, roomID string) ([]entity.Chat, error)
	GetRecents(ctx context.Context, tx *sql.Tx, username string) ([]model.RecentChat, error)
	Delete(ctx context.Context, tx *sql.Tx, chatID int) error
}
