package service

import (
	"context"
	"database/sql"
	"strings"

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

// Save stores new message into database.
func (s *chatService) Save(ctx context.Context, chat entity.Chat) error {
	chat.Username = strings.TrimSpace(chat.Username)
	if int8(len(chat.RoomID)) != entity.RoomIDLength {
		return &RoomErrInvalidRoomID{}
	}
	exists, err := s.occupantRepo.Exists(ctx, chat.RoomID, chat.Username)
	if err != nil {
		return err
	}
	if !exists {
		return sql.ErrNoRows
	}
	return s.chatRepo.Save(ctx, chat)
}

func (s *chatService) GetAll(ctx context.Context, roomID string) ([]entity.Chat, error) {
	// wip
}

func (s *chatService) Delete(ctx context.Context, chatID int) error {
	// wip
}
