package handler

type PathErrNoPathValue struct{}

func (e *PathErrNoPathValue) Error() string {
	return "error path value is empty or missing"
}
