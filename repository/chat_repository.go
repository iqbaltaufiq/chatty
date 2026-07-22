package repository

import (
	"context"

	"chatty/entity"
)

type ChatRepository interface {
	Save(ctx context.Context, chat entity.Chat) (entity.Chat, error)
	GetAll(ctx context.Context, roomID string) ([]entity.Chat, error)
	GetRecents(ctx context.Context, username string) ([]entity.Chat, error)
	Delete(ctx context.Context, chatID int) error
}
