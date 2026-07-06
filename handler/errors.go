package handler

type PathErrNoPathValue struct{}

func (e *PathErrNoPathValue) Error() string {
	return "error path value is empty or missing"
}

type PathErrInvalidPath struct{}

func (e *PathErrInvalidPath) Error() string {
	return "error path value does not meet expectation"
}
