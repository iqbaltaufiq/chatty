package customerrors

type ErrInvalidFileFormat struct{}

func (e *ErrInvalidFileFormat) Error() string {
	return "unsupported file format"
}
