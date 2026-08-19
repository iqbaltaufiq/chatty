package service

import (
	"context"
	"io"

	"chatty/repository"
)

type fileService struct {
	fileRepo repository.FileRepository
}

func NewFileService(fileRepo repository.FileRepository) FileService {
	return &fileService{fileRepo}
}

func (s *fileService) Store(ctx context.Context, file io.Reader, filename string) error {
	return s.fileRepo.Store(ctx, file, filename)
}
