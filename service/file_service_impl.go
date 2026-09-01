package service

import (
	"context"
	"image"
	"io"
	"os"
	"path/filepath"

	"chatty/repository"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

type fileService struct {
	fileRepo repository.FileRepository
}

func NewFileService(fileRepo repository.FileRepository) FileService {
	return &fileService{fileRepo}
}

func (s *fileService) Store(ctx context.Context, file io.Reader, filename string) error {
	return s.fileRepo.Store(ctx, file, filename)
}

func (s *fileService) StoreAvatar(ctx context.Context, src io.Reader, filename string) error {
	dest := filepath.Join("public", "static", "assets", filename)
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	return makeSquareAvatar(src, f, 400)
}

func (s *fileService) Delete(ctx context.Context, filename string) error {
	target := filepath.Join("public", "static", "assets", filename)
	return os.Remove(target)
}
