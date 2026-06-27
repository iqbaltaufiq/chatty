package repository

import (
	"context"
	"database/sql"

	"chatty/entity"
)

type chatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) ChatRepository {
	return &chatRepository{db}
}

// Save stores a message into database.
func (r *chatRepository) Save(ctx context.Context, chat entity.Chat) error {
	query := `
	INSERT INTO chats (username, message, room_id)
	VALUES (?,?,?)
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	_, err := stmt.ExecContext(
		ctx,
		chat.Username,
		chat.Message,
		chat.RoomID,
	)
	return err
}

// Delete removes a message from database.
func (r *chatRepository) Delete(ctx context.Context, chatID int) error {
	// not yet implemented
	return nil
}

// GetAll retrieves all the messages that belongs to a room.
func (r *chatRepository) GetAll(ctx context.Context, roomID string) ([]entity.Chat, error) {
	// not yet implemented
	return nil, nil
}
