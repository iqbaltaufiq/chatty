package model

import (
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	customerrors "chatty/model/custom_errors"
	"chatty/util"
)

const maxFileSize = 2 * 1000 * 1000 // 2MB
var allowedExt = map[string]struct{}{
	"jpg":  {},
	"jpeg": {},
	"png":  {},
	"webp": {},
	"avif": {},
	"tiff": {},
}

type UploadedPhoto struct {
	File   io.Reader
	Header *multipart.FileHeader
}

func (up *UploadedPhoto) IsValid() error {
	if up.File == nil || up.Header == nil {
		return http.ErrMissingFile
	}
	if up.Header.Size > maxFileSize {
		return &customerrors.ErrSizeLimitExceeded{}
	}
	s := strings.Split(up.Header.Filename, ".")
	if _, exists := allowedExt[s[len(s)-1]]; !exists {
		return &customerrors.ErrInvalidFileFormat{}
	}
	return nil
}

func (up *UploadedPhoto) GetFilename() string {
	return up.Header.Filename
}

func (up *UploadedPhoto) HashFilename() string {
	if up.IsValid() != nil {
		return ""
	}
	hash, _ := util.HashFilename(up.Header.Filename)
	return hash
}
