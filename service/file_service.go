package service

import (
	"context"
	"io"

	"chatty/model"
)

type FileService interface {
	Store(ctx context.Context, file io.Reader, filename string) error
	StoreAvatar(ctx context.Context, photo model.UploadedPhoto) error
}
