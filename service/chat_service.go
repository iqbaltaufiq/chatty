package service

import (
	"context"

	"chatty/entity"
	"chatty/model"
)

type ChatService interface {
	Save(ctx context.Context, chat entity.Chat) (entity.Chat, error)
	GetAll(ctx context.Context, roomID string) ([]entity.Chat, error)
	GetRecents(ctx context.Context) ([]model.RecentChat, error)
	Delete(ctx context.Context, chatID int) error
	DeleteAll(ctx context.Context, roomID string) error
}
