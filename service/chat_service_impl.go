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

// GetRecents retrieves the most recent messages from every chatroom this user belongs to.
// The user is taken from the authenticated user in the request context.
func (s *chatService) GetRecents(ctx context.Context) ([]entity.Chat, error) {
	username, ok := util.FromCtxUsername(ctx)
	if !ok {
		return nil, &util.ContextErrNotFound{}
	}
	return s.chatRepo.GetRecents(ctx, username)
}

// Delete deletes a chat specified by chatID.
func (s *chatService) Delete(ctx context.Context, chatID int) error {
	return s.chatRepo.Delete(ctx, chatID)
}
