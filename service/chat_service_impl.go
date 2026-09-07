package service

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"chatty/entity"
	"chatty/model"
	customerrors "chatty/model/custom_errors"
	"chatty/repository"
	"chatty/util"
)

type chatService struct {
	db           *sql.DB
	chatRepo     repository.ChatRepository
	occupantRepo repository.OccupantRepository
}

func NewChatService(
	db *sql.DB,
	chatRepo repository.ChatRepository,
	occupantRepo repository.OccupantRepository,
) ChatService {
	return &chatService{
		db,
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
	tx, _ := s.db.Begin()
	defer tx.Rollback()

	exists, err := s.occupantRepo.Exists(ctx, tx, chat.RoomID, chat.Username)
	if err != nil {
		return entity.Chat{}, err
	}
	if !exists {
		return entity.Chat{}, sql.ErrNoRows
	}
	timeNowUTC := time.Now().UTC()
	chat.CreatedAt = &timeNowUTC
	chat.UpdatedAt = &timeNowUTC

	chat, err = s.chatRepo.Save(ctx, tx, chat)
	if err != nil {
		return entity.Chat{}, err
	}
	return chat, tx.Commit()
}

// GetAll retrieves all of the chats belong to the specified room.
func (s *chatService) GetAll(ctx context.Context, roomID string) ([]entity.Chat, error) {
	if int8(len(roomID)) != entity.RoomIDLength {
		return nil, &RoomErrInvalidRoomID{}
	}
	username, _ := util.FromCtxUsername(ctx)

	tx, _ := s.db.Begin()
	defer tx.Rollback()

	exists, err := s.occupantRepo.Exists(ctx, tx, roomID, username)
	if err != nil || !exists {
		return nil, sql.ErrNoRows
	}
	chats, err := s.chatRepo.GetAll(ctx, tx, roomID)
	if err != nil {
		return nil, err
	}
	return chats, tx.Commit()
}

// GetRecents retrieves the most recent messages from every chatroom this user belongs to.
// The user is taken from the authenticated user in the request context.
func (s *chatService) GetRecents(ctx context.Context) ([]model.RecentChat, error) {
	username, ok := util.FromCtxUsername(ctx)
	if !ok {
		return nil, &util.ContextErrNotFound{}
	}
	return s.chatRepo.GetRecents(ctx, s.db, username)
}

// Delete deletes a chat specified by chatID.
func (s *chatService) Delete(ctx context.Context, chatID int) error {
	return s.chatRepo.Delete(ctx, s.db, chatID)
}

func (s *chatService) DeleteAll(ctx context.Context, roomID string) error {
	if len(roomID) != int(entity.RoomIDLength) {
		return &customerrors.ErrInvalidPath{}
	}
	username, _ := util.FromCtxUsername(ctx)

	tx, _ := s.db.Begin()
	defer tx.Rollback()

	exists, err := s.occupantRepo.Exists(ctx, tx, roomID, username)
	if err != nil {
		return err
	}
	if !exists {
		return sql.ErrNoRows
	}
	if err := s.chatRepo.DeleteAll(ctx, tx, roomID); err != nil {
		return err
	}
	return tx.Commit()
}
