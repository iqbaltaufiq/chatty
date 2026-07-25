package repository

import (
	"context"

	"chatty/entity"
	"chatty/model"
)

type ChatRepository interface {
	Save(ctx context.Context, chat entity.Chat) (entity.Chat, error)
	GetAll(ctx context.Context, roomID string) ([]entity.Chat, error)
	GetRecents(ctx context.Context, username string) ([]model.RecentChat, error)
	Delete(ctx context.Context, chatID int) error
}
