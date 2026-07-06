package handler

type PathErrNoPathValue struct{}

func (e *PathErrNoPathValue) Error() string {
	return "path value is empty or missing"
}

type PathErrInvalidPath struct{}

func (e *PathErrInvalidPath) Error() string {
	return "path value does not meet expectation"
}
