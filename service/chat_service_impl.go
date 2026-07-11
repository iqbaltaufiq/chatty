package service

import (
	"context"

	"chatty/entity"
	"chatty/repository"
)

type chatService struct {
	chatRepo     repository.ChatRepository
	occupantRepo repository.OccupantRepository
}

func NewChatService(
	chatRepo repository.ChatRepository,
	occupantRepo repository.OccupantRepository,
) ChatService {
	return &chatService{
		chatRepo,
		occupantRepo,
	}
}

func (s *chatService) Save(ctx context.Context, chat entity.Chat) error {
	// wip
}

func (s *chatService) GetAll(ctx context.Context, roomID string) ([]entity.Chat, error) {
	// wip
}

func (s *chatService) Delete(ctx context.Context, chatID int) error {
	// wip
}
