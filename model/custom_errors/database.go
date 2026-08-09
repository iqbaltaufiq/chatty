package customerrors

type RowsErrAlreadyExists struct{}

func (e *RowsErrAlreadyExists) Error() string {
	return "expected to be absent, but resource already exists in database"
}
