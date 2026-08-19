package repository

import (
	"context"
	"io"
)

type FileRepository interface {
	Store(ctx context.Context, file io.Reader, filename string) error
	Delete(ctx context.Context, filename string) error
}
