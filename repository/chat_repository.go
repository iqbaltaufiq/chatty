package repository

import (
	"context"

	"chatty/entity"
	"chatty/model"
)

type ChatRepository interface {
	Save(ctx context.Context, dbtx DBTX, chat entity.Chat) (entity.Chat, error)
	GetAll(ctx context.Context, dbtx DBTX, roomID string) ([]entity.Chat, error)
	GetRecents(ctx context.Context, dbtx DBTX, username string) ([]model.RecentChat, error)
	Delete(ctx context.Context, dbtx DBTX, chatID int) error
}
