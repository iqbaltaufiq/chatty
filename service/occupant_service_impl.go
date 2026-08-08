package service

import (
	"context"
	"database/sql"
	"strings"

	"chatty/repository"
)

type occupantService struct {
	db           *sql.DB
	occupantRepo repository.OccupantRepository
}

func NewOccupantService(
	db *sql.DB,
	occupantRepo repository.OccupantRepository,
) OccupantService {
	return &occupantService{
		db,
		occupantRepo,
	}
}

// Add checks if the user has not occupied the speficied room
// and add them into the room.
func (s *occupantService) Add(ctx context.Context, roomID, username string) error {
	roomID = strings.TrimSpace(roomID)
	username = strings.TrimSpace(username)

	tx, _ := s.db.Begin()
	defer tx.Rollback()

	exists, err := s.occupantRepo.Exists(ctx, tx, roomID, username)
	if err != nil {
		return err
	}
	if exists {
		return &RowsErrAlreadyExists{}
	}
	if err := s.occupantRepo.Add(ctx, tx, roomID, username); err != nil {
		return err
	}
	return tx.Commit()
}

// Remove checks if the user belongs to the specified room
// and remove them from the room.
func (s *occupantService) Remove(ctx context.Context, roomID, username string) error {
	roomID = strings.TrimSpace(roomID)
	username = strings.TrimSpace(username)

	tx, _ := s.db.Begin()
	defer tx.Rollback()

	exists, err := s.occupantRepo.Exists(ctx, tx, roomID, username)
	if err != nil || !exists {
		return sql.ErrNoRows
	}
	if err := s.occupantRepo.Remove(ctx, tx, roomID, username); err != nil {
		return err
	}
	return tx.Commit()
}
