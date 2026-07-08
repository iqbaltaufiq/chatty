package service

type fileExtension map[string]struct{}

func (f fileExtension) Exists(ext string) bool {
	_, ok := f[ext]
	return ok
}

// imageExtensions contain the allowed file format for images.
var imageExtensions = fileExtension{
	"jpg":  {},
	"jpeg": {},
	"png":  {},
	"webp": {},
	"avif": {},
	"tiff": {},
}
