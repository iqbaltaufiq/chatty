package service

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"chatty/entity"
	"chatty/repository"
	"chatty/util"
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
func (s *chatService) Save(ctx context.Context, chat entity.Chat) (entity.Chat, error) {
	chat.Username = strings.TrimSpace(chat.Username)
	if int8(len(chat.RoomID)) != entity.RoomIDLength {
		return entity.Chat{}, &RoomErrInvalidRoomID{}
	}
	exists, err := s.occupantRepo.Exists(ctx, chat.RoomID, chat.Username)
	if err != nil {
		return entity.Chat{}, err
	}
	if !exists {
		return entity.Chat{}, sql.ErrNoRows
	}
	timeNowUTC := time.Now().UTC()
	chat.CreatedAt = &timeNowUTC
	chat.UpdatedAt = &timeNowUTC

	return s.chatRepo.Save(ctx, chat)
}

// GetAll retrieves all of the chats belong to the specified room.
func (s *chatService) GetAll(ctx context.Context, roomID string) ([]entity.Chat, error) {
	if int8(len(roomID)) != entity.RoomIDLength {
		return nil, &RoomErrInvalidRoomID{}
	}
	username, _ := util.FromCtxUsername(ctx)

	exists, err := s.occupantRepo.Exists(ctx, roomID, username)
	if err != nil || !exists {
		return nil, sql.ErrNoRows
	}
	return s.chatRepo.GetAll(ctx, roomID)
}

// Delete deletes a chat specified by chatID.
func (s *chatService) Delete(ctx context.Context, chatID int) error {
	return s.chatRepo.Delete(ctx, chatID)
}
