package service

import (
	"context"

	"chatty/repository"
)

type occupantService struct {
	occupantRepo repository.OccupantRepository
}

func NewOccupantService(occupantRepo repository.OccupantRepository) OccupantService {
	return &occupantService{occupantRepo}
}

func (s *occupantService) Add(ctx context.Context, roomID, username string) error {
	// not yet implemented
}

func (s *occupantService) Remove(ctx context.Context, roomID, username string) error {
	// not yet implemented
}
