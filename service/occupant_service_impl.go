package service

import (
	"context"
	"database/sql"
	"strings"

	"chatty/repository"
)

type occupantService struct {
	occupantRepo repository.OccupantRepository
}

func NewOccupantService(occupantRepo repository.OccupantRepository) OccupantService {
	return &occupantService{occupantRepo}
}

func (s *occupantService) Add(ctx context.Context, roomID, username string) error {
	roomID = strings.TrimSpace(roomID)
	username = strings.TrimSpace(username)
	exists, err := s.occupantRepo.Exists(ctx, roomID, username)
	if err != nil {
		return err
	}
	if exists {
		return &RowsErrAlreadyExists{}
	}
	return s.occupantRepo.Add(ctx, roomID, username)
}

func (s *occupantService) Remove(ctx context.Context, roomID, username string) error {
	roomID = strings.TrimSpace(roomID)
	username = strings.TrimSpace(username)
	exists, err := s.occupantRepo.Exists(ctx, roomID, username)
	if err != nil || !exists {
		return sql.ErrNoRows
	}
	return s.occupantRepo.Remove(ctx, roomID, username)
}
