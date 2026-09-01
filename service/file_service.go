package service

import (
	"context"
	"io"
)

type FileService interface {
	Store(ctx context.Context, file io.Reader, filename string) error
	StoreAvatar(ctx context.Context, src io.Reader, filename string) error
	Delete(ctx context.Context, filename string) error
}
