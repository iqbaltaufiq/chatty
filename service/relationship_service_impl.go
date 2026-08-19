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

type relationshipService struct {
	db               *sql.DB
	relationshipRepo repository.RelationshipRepository
	roomRepo         repository.RoomRepository
	occupantRepo     repository.OccupantRepository
}

func NewRelationshipService(
	db *sql.DB,
	relationshipRepo repository.RelationshipRepository,
	roomRepo repository.RoomRepository,
	occupantRepo repository.OccupantRepository,
) RelationshipService {
	return &relationshipService{
		db,
		relationshipRepo,
		roomRepo,
		occupantRepo,
	}
}

// Add befriends initiator with the target.
//
// Returns error if they are already friends.
func (s *relationshipService) Add(ctx context.Context, initiator, target string) error {
	initiator = strings.TrimSpace(initiator)
	target = strings.TrimSpace(target)

	tx, _ := s.db.Begin()
	defer tx.Rollback()

	exists, err := s.relationshipRepo.Exists(ctx, tx, initiator, target)
	if err != nil {
		return err
	}
	if exists {
		return &customerrors.RowsErrAlreadyExists{}
	}
	if err := s.relationshipRepo.Add(ctx, tx, initiator, target); err != nil {
		return err
	}
	roomID := ""
	for {
		roomID = util.GenerateAlphanumeric(int(entity.RoomIDLength))
		exists, err := s.roomRepo.Exists(ctx, tx, roomID)
		if err != nil {
			return err
		}
		if !exists {
			break
		}
	}
	current := time.Now().UTC()
	s.roomRepo.Create(ctx, tx, entity.Room{
		ID:        roomID,
		Name:      "New Chatroom",
		IsPrivate: true,
		CreatedAt: &current,
	})
	if err := s.occupantRepo.Add(ctx, tx, roomID, initiator); err != nil {
		return err
	}
	if err := s.occupantRepo.Add(ctx, tx, roomID, target); err != nil {
		return err
	}
	return tx.Commit()
}

// Remove unfriends both users.
// It deletes the relationship and the 1-on-1 chatroom including all the chats.
//
// Returns error if they haven't become friends yet.
func (s *relationshipService) Remove(ctx context.Context, initiator, target string) error {
	initiator = strings.TrimSpace(initiator)
	target = strings.TrimSpace(target)

	tx, _ := s.db.Begin()
	defer tx.Rollback()

	if err := s.relationshipRepo.Remove(ctx, tx, initiator, target); err != nil {
		return err
	}
	roomID, _ := s.occupantRepo.GetRoomBetweenUsers(ctx, tx, initiator, target)
	if err := s.roomRepo.Delete(ctx, tx, roomID); err != nil {
		return err
	}
	return tx.Commit()
}

// RemoveByID unfriends two users referenced by their relationshipID.
func (s *relationshipService) RemoveByID(ctx context.Context, relationshipID int) error {
	tx, _ := s.db.Begin()
	defer tx.Rollback()

	rel, err := s.relationshipRepo.Get(ctx, tx, relationshipID)
	if err != nil {
		return err
	}
	roomID, err := s.occupantRepo.GetRoomBetweenUsers(ctx, tx, rel.Initiator, rel.Target)
	if err != nil {
		return err
	}
	if err := s.relationshipRepo.RemoveByID(ctx, s.db, relationshipID); err != nil {
		return err
	}
	if err := s.roomRepo.Delete(ctx, tx, roomID); err != nil {
		return err
	}
	return tx.Commit()
}

// GetFriends retrieves a list of friends that belongs to the user.
func (s *relationshipService) GetFriends(ctx context.Context, username string) ([]model.Friend, error) {
	username = strings.TrimSpace(username)
	return s.relationshipRepo.GetFriends(ctx, s.db, username)
}
