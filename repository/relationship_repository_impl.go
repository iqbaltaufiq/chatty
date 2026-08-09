package repository

import (
	"context"
	"database/sql"

	"chatty/model"
	"chatty/util"
)

type relationshipRepository struct{}

func NewRelationshipRepository() RelationshipRepository {
	return &relationshipRepository{}
}

// Add befriends initiator with the target.
func (r *relationshipRepository) Add(ctx context.Context, dbtx DBTX, initiator, target string) error {
	query := `
	INSERT INTO relationships (initiator, target)
	VALUES (?,?)
	`
	_, err := dbtx.ExecContext(ctx, query, initiator, target)
	return err
}

func (r *relationshipRepository) Remove(ctx context.Context, dbtx DBTX, self, target string) error {
	query := `
	DELETE FROM relationships
	WHERE initiator = ? AND target = ? OR initiator = ? AND target = ?
	LIMIT 1
	`
	args := [4]any{self, target, target, self}
	_, err := dbtx.ExecContext(ctx, query, args[:]...)
	return err
}

// Remove removes relationship between two users.
func (r *relationshipRepository) RemoveByID(ctx context.Context, dbtx DBTX, relationshipID int) error {
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
	if err := dbtx.QueryRowContext(ctx, fetchQuery, fetchArgs...).
		Scan(&rowCount); err != nil {
		return err
	}
	if rowCount == 0 {
		return &RelationshipErrDoesNotBelong{}
	}
	if _, err := dbtx.ExecContext(ctx, deleteQuery, relationshipID); err != nil {
		return err
	}
	return nil
}

// Exists checks if two users have befriended with each other.
func (r *relationshipRepository) Exists(ctx context.Context, dbtx DBTX, initiator, target string) (bool, error) {
	query := `
	SELECT COUNT(*) as 'rowsCount'
	FROM relationships
	WHERE (initiator = ? AND target = ?) OR (initiator = ? AND target = ?)
	LIMIT 1
	`
	args := []any{initiator, target, target, initiator}
	row := dbtx.QueryRowContext(ctx, query, args...)
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
func (r *relationshipRepository) GetFriends(ctx context.Context, dbtx DBTX, username string) ([]model.Friend, error) {
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
	args := [3]any{username, username, username}
	rows, err := dbtx.QueryContext(ctx, query, args[:]...)
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
