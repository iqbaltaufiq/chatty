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
func (r *relationshipRepository) Add(ctx context.Context, initiator, target string) error {
	query := `
	INSERT INTO relationships (initiator, target)
	VALUES (?,?)
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	_, err := stmt.ExecContext(ctx, initiator, target)
	return err
}

func (r *relationshipRepository) Remove(ctx context.Context, self, target string) error {
	query := `
	DELETE FROM relationships
	WHERE initiator = ? AND target = ? OR initiator = ? AND target = ?
	LIMIT 1
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	_, err := stmt.ExecContext(
		ctx,
		self, target,
		target, self,
	)
	return err
}

// Remove removes relationship between two users.
func (r *relationshipRepository) RemoveByID(ctx context.Context, relationshipID int) error {
	query := `
	DELETE FROM relationships
	WHERE id = ?
	LIMIT 1
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	_, err := stmt.ExecContext(ctx, relationshipID)
	return err
}

// Exists checks if two users have befriended with each other.
func (r *relationshipRepository) Exists(ctx context.Context, initiator, target string) (bool, error) {
	query := `
	SELECT COUNT(*) as 'rowsCount'
	FROM relationships
	WHERE (initiator = ? AND target = ?) OR (initiator = ? AND target = ?)
	LIMIT 1
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	args := []any{initiator, target, target, initiator}
	row := stmt.QueryRowContext(ctx, args...)

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
func (r *relationshipRepository) GetFriends(ctx context.Context, username string) ([]model.Friend, error) {
	query := `
	SELECT source.id, u.uuid, u.username, u.fullname, u.photo
	FROM (
		SELECT r.id, r.initiator as 'friend'
		FROM relationships r
		WHERE r.target = ?
		UNION
		SELECT r.id, r.target as 'friend'
		FROM relationships r
		WHERE r.initiator = ?
	) AS source
	JOIN users AS u ON u.username = source.friend
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, username, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []model.Friend
	for rows.Next() {
		var friend model.Friend
		var binaryUUID []byte

		if err := rows.Scan(
			&friend.RelationshipID,
			&binaryUUID,
			&friend.Username,
			&friend.Fullname,
			&friend.Photo,
		); err != nil {
			return nil, err
		}
		friend.UUID, _ = util.UUIDtoString(binaryUUID)
		friends = append(friends, friend)
	}
	return friends, nil
}
