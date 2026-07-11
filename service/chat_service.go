package service

import (
	"context"

	"chatty/entity"
)

type ChatService interface {
	Save(ctx context.Context, chat entity.Chat) error
	GetAll(ctx context.Context, roomID string) ([]entity.Chat, error)
	Delete(ctx context.Context, chatID int) error
}
