package model

import (
	"io"
	"mime/multipart"
	"net/http"

	customerrors "chatty/model/custom_errors"
)

const maxFileSize = 2 * 1000 * 1000 // 2MB

type UploadedPhoto struct {
	File     io.Reader
	Header   *multipart.FileHeader
	Username string
	OldPhoto string
}

func (up *UploadedPhoto) IsValid() error {
	if up.File == nil || up.Header == nil {
		return http.ErrMissingFile
	}
	if up.Username == "" || up.OldPhoto == "" {
		return &customerrors.ErrInsufficientPayload{}
	}
	if up.Header.Size > maxFileSize {
		return &customerrors.ErrSizeLimitExceeded{}
	}
	return nil
}

func (up *UploadedPhoto) GetFilename() string {
	return up.Header.Filename
}
