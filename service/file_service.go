package service

import (
	"context"
	"io"
)

type FileService interface {
	Store(ctx context.Context, file io.Reader, filename string) error
}
