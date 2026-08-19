package repository

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

type fileRepository struct{}

func NewFileRepository() FileRepository {
	return &fileRepository{}
}

func (r *fileRepository) Store(ctx context.Context, file io.Reader, filename string) error {
	dest := filepath.Join("public", "static", "assets", filename)
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	return err
}

func (r *fileRepository) Delete(ctx context.Context, filename string) error {
	target := filepath.Join("public", "static", "assets", filename)
	return os.Remove(target)
}
