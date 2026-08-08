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

	tx, exists, err := s.occupantRepo.Exists(ctx, tx, roomID, username)
	if err != nil {
		return err
	}
	if exists {
		return &RowsErrAlreadyExists{}
	}
	tx, err = s.occupantRepo.Add(ctx, tx, roomID, username)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// Remove checks if the user belongs to the specified room
// and remove them from the room.
func (s *occupantService) Remove(ctx context.Context, roomID, username string) error {
	roomID = strings.TrimSpace(roomID)
	username = strings.TrimSpace(username)
	exists, err := s.occupantRepo.Exists(ctx, roomID, username)
	if err != nil || !exists {
		return sql.ErrNoRows
	}
	return s.occupantRepo.Remove(ctx, roomID, username)
}
