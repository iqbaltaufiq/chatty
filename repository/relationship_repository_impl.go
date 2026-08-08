package repository

import (
	"context"
	"database/sql"

	"chatty/model"
	"chatty/util"
)

type relationshipRepository struct {
	db *sql.DB
}

func NewRelationshipRepository(db *sql.DB) RelationshipRepository {
	return &relationshipRepository{db}
}

// Add befriends initiator with the target.
func (r *relationshipRepository) Add(ctx context.Context, tx *sql.Tx, initiator, target string) error {
	query := `
	INSERT INTO relationships (initiator, target)
	VALUES (?,?)
	`
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, initiator, target)
	} else {
		_, err = r.db.ExecContext(ctx, query, initiator, target)
	}
	return err
}

func (r *relationshipRepository) Remove(ctx context.Context, tx *sql.Tx, self, target string) error {
	query := `
	DELETE FROM relationships
	WHERE initiator = ? AND target = ? OR initiator = ? AND target = ?
	LIMIT 1
	`
	var err error
	args := [4]any{self, target, target, self}
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, args[:]...)
	} else {
		_, err = r.db.ExecContext(ctx, query, args[:]...)
	}

	return err
}

// Remove removes relationship between two users.
func (r *relationshipRepository) RemoveByID(ctx context.Context, tx *sql.Tx, relationshipID int) error {
	fetchQuery := `
	SELECT COUNT(*) AS 'rowCount'
	FROM relationships
	WHERE id = ? AND (initiator = ? OR target = ?)
	LIMIT 1
	`
	deleteQuery := `
	DELETE FROM relationships
	WHERE id = ?
	LIMIT 1
	`
	var (
		self, _   = util.FromCtxUsername(ctx)
		rowCount  = 0
		fetchArgs = []any{relationshipID, self, self}
	)
	if err := tx.QueryRowContext(ctx, fetchQuery, fetchArgs...).
		Scan(&rowCount); err != nil {
		return err
	}
	if rowCount == 0 {
		return &RelationshipErrDoesNotBelong{}
	}
	if _, err := tx.ExecContext(ctx, deleteQuery, relationshipID); err != nil {
		return err
	}
	return nil
}

// Exists checks if two users have befriended with each other.
func (r *relationshipRepository) Exists(ctx context.Context, tx *sql.Tx, initiator, target string) (bool, error) {
	query := `
	SELECT COUNT(*) as 'rowsCount'
	FROM relationships
	WHERE (initiator = ? AND target = ?) OR (initiator = ? AND target = ?)
	LIMIT 1
	`
	row := new(sql.Row)
	args := []any{initiator, target, target, initiator}
	if tx != nil {
		row = tx.QueryRowContext(ctx, query, args...)
	} else {
		row = r.db.QueryRowContext(ctx, query, args...)
	}
	exists := false
	rowsCount := 0
	if err := row.Scan(&rowsCount); err != nil {
		return exists, err
	}
	if rowsCount > 0 {
		exists = true
	}
	return exists, nil
}

// GetFriends retrieves a list of users that have relationship with the specified user.
func (r *relationshipRepository) GetFriends(ctx context.Context, tx *sql.Tx, username string) ([]model.Friend, error) {
	query := `
	SELECT source.id, o.room_id, u.uuid, source.username, u.fullname, u.photo
	FROM (
		SELECT r.id, r.target AS username
		FROM relationships r
		WHERE r.initiator = ?
		UNION
		SELECT r.id, r.initiator AS username
		FROM relationships r
		WHERE r.target = ?
	) AS source
	JOIN users u ON u.username = source.username
	LEFT JOIN occupants o ON o.username = source.username AND o.room_id IN (SELECT room_id FROM occupants WHERE username = ?)
	ORDER BY u.fullname ASC
	`
	var err error
	rows := new(sql.Rows)
	args := [3]any{username, username, username}

	if tx != nil {
		rows, err = tx.QueryContext(ctx, query, args[:]...)
	} else {
		rows, err = r.db.QueryContext(ctx, query, args[:]...)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []model.Friend
	for rows.Next() {
		var (
			friend         model.Friend
			binaryUUID     []byte
			nullableRoomID sql.NullString
		)
		if err := rows.Scan(
			&friend.RelationshipID,
			&nullableRoomID,
			&binaryUUID,
			&friend.Username,
			&friend.Fullname,
			&friend.Photo,
		); err != nil {
			return nil, err
		}
		friend.RoomID = ""
		if nullableRoomID.Valid {
			friend.RoomID = nullableRoomID.String
		}
		friend.UUID, _ = util.UUIDtoString(binaryUUID)
		friends = append(friends, friend)
	}
	return friends, nil
}
