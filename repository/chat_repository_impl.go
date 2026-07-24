package repository

import (
	"context"
	"database/sql"

	"chatty/entity"
	"chatty/util"
)

type chatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) ChatRepository {
	return &chatRepository{db}
}

// Save stores a message into database.
func (r *chatRepository) Save(ctx context.Context, chat entity.Chat) (entity.Chat, error) {
	query := `
	INSERT INTO chats (username, room_id, message, created_at, updated_at)
	VALUES (?,?,?,?,?)
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	row, err := stmt.ExecContext(
		ctx,
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
func (r *chatRepository) Delete(ctx context.Context, chatID int) error {
	query := `
	DELETE FROM chats
	WHERE id = ? AND username = ?
	LIMIT 1
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	username, _ := util.FromCtxUsername(ctx)
	row, err := stmt.ExecContext(ctx, chatID, username)
	if err != nil {
		return err
	}
	if rowsAffected, _ := row.RowsAffected(); rowsAffected < 1 {
		return sql.ErrNoRows
	}
	return err
}

// GetAll retrieves all the messages that belongs to a room.
func (r *chatRepository) GetAll(ctx context.Context, roomID string) ([]entity.Chat, error) {
	const MaxRows int = 25
	query := `
	SELECT id, username, message, room_id, created_at, updated_at
	FROM chats
	WHERE room_id = ?
	LIMIT ?
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, roomID, MaxRows)
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
func (r *chatRepository) GetRecents(ctx context.Context, username string) ([]entity.Chat, error) {
	query := `
	SELECT ch.id, ch.username, ch.room_id, ch.message, ch.created_at
	FROM chats ch
	JOIN (
		SELECT max(id) AS max_id
		FROM chats c
		JOIN (
			SELECT room_id FROM occupants
			WHERE username = ?
		) AS filteredID ON filteredID.room_id = c.room_id
		GROUP BY c.room_id
	) AS maxid ON maxid.max_id = ch.id
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chats := []entity.Chat{}
	for rows.Next() {
		chat := entity.Chat{}
		if err := rows.Scan(
			&chat.ID,
			&chat.Username,
			&chat.RoomID,
			&chat.Message,
			&chat.CreatedAt,
		); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}
