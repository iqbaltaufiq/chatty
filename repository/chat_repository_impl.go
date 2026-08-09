package repository

import (
	"context"
	"database/sql"

	"chatty/entity"
	"chatty/model"
	"chatty/util"
)

type chatRepository struct{}

func NewChatRepository() ChatRepository {
	return &chatRepository{}
}

// Save stores a message into database.
func (r *chatRepository) Save(ctx context.Context, dbtx DBTX, chat entity.Chat) (entity.Chat, error) {
	query := `
	INSERT INTO chats (username, room_id, message, created_at, updated_at)
	VALUES (?,?,?,?,?)
	`
	row, err := dbtx.ExecContext(
		ctx,
		query,
		chat.Username,
		chat.RoomID,
		chat.Message,
		chat.CreatedAt,
		chat.UpdatedAt,
	)
	if err != nil {
		return entity.Chat{}, err
	}
	chat.ID, _ = row.LastInsertId()

	return chat, err
}

// Delete removes a message from database.
func (r *chatRepository) Delete(ctx context.Context, dbtx DBTX, chatID int) error {
	query := `
	DELETE FROM chats
	WHERE id = ? AND username = ?
	LIMIT 1
	`
	username, _ := util.FromCtxUsername(ctx)
	row, err := dbtx.ExecContext(ctx, query, chatID, username)
	if err != nil {
		return err
	}
	if rowsAffected, _ := row.RowsAffected(); rowsAffected < 1 {
		return sql.ErrNoRows
	}
	return err
}

// GetAll retrieves all the messages that belongs to a room.
func (r *chatRepository) GetAll(ctx context.Context, dbtx DBTX, roomID string) ([]entity.Chat, error) {
	const MaxRows int = 25
	query := `
	SELECT id, username, message, room_id, created_at, updated_at
	FROM chats
	WHERE room_id = ?
	LIMIT ?
	`
	rows, err := dbtx.QueryContext(ctx, query, roomID, MaxRows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chats := make([]entity.Chat, 0, MaxRows)
	for rows.Next() {
		var chat entity.Chat
		if err := rows.Scan(
			&chat.ID,
			&chat.Username,
			&chat.Message,
			&chat.RoomID,
			&chat.CreatedAt,
			&chat.UpdatedAt,
		); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}

// GetRecents retrieves the most recent messages from every room this user belongs to.
func (r *chatRepository) GetRecents(ctx context.Context, dbtx DBTX, username string) ([]model.RecentChat, error) {
	query := `
	SELECT 
		ch.id,
		o_exc.username,
		u.fullname,
		ch.room_id,
		u.photo,
		ch.message,
		ch.created_at
	FROM (
		SELECT max(c.id) AS max_id, c.room_id
		FROM chats c
		JOIN occupants oc ON oc.room_id = c.room_id
		WHERE oc.username = ?
		GROUP BY c.room_id
	) AS newest
	JOIN chats ch ON ch.id = newest.max_id
	JOIN occupants o_exc ON o_exc.room_id = newest.room_id AND o_exc.username != ?
	JOIN users u ON u.username = o_exc.username
	ORDER BY ch.id DESC
	`
	rows, err := dbtx.QueryContext(ctx, query, username, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chats := []model.RecentChat{}
	for rows.Next() {
		chat := model.RecentChat{}
		if err := rows.Scan(
			&chat.ID,
			&chat.Username,
			&chat.Fullname,
			&chat.RoomID,
			&chat.Photo,
			&chat.Message,
			&chat.CreatedAt,
		); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}
