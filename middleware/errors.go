package middleware

type AuthErrAlreadyLoggedIn struct{}

func (e *AuthErrAlreadyLoggedIn) Error() string {
	return "request is coming from an already authenticated user"
}
