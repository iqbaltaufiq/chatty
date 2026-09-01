package service

import (
	"context"
	"image"
	"io"
	"os"
	"path/filepath"

	"chatty/model"
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

func (s *fileService) StoreAvatar(ctx context.Context, photo model.UploadedPhoto) error {
	if err := photo.IsValid(); err != nil {
		return err
	}
	newFilename := filepath.Join("public", "static", "assets", photo.GetFilename())
	f, err := os.Create(newFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	return makeSquareAvatar(photo.File, f)
}
