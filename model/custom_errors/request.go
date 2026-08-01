package customerrors

type ErrNoPathValue struct{}

func (e *ErrNoPathValue) Error() string {
	return "path value is empty or missing"
}

type ErrInvalidPath struct{}

func (e *ErrInvalidPath) Error() string {
	return "path value does not meet expectation"
}

type ErrInsufficientPayload struct{}

func (e *ErrInsufficientPayload) Error() string {
	return "request payload does not meet the required expectation"
}

type ErrSizeLimitExceeded struct{}

func (e *ErrSizeLimitExceeded) Error() string {
	return "payload exceeded the maximum size allowed"
}
