package util

import "context"

// contextKey is the custom type for key context.
type contextKey struct{}

var (
	usernameKey contextKey
	uuidKey     contextKey
)

// FromCtxUsername takes out username from the context.
//
// This getter is necessary since context key is non-exportable.
func FromCtxUsername(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(usernameKey).(string)
	return username, ok
}

// FromCtxUUID takes out userID (uuid) from the context.
//
// This getter is necessary since context key is non-exportable.
func FromCtxUUID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(uuidKey).(string)
	return userID, ok
}

// SetCtxUsername attaches `username` into the `parent` context.
// Returns the appended context.
func SetCtxUsername(parent context.Context, username string) context.Context {
	return context.WithValue(parent, usernameKey, username)
}

// SetCtxUUID attaches `uuid` into the `parent` context.
// Returns the appended context.
func SetCtxUUID(parent context.Context, uuid string) context.Context {
	return context.WithValue(parent, uuidKey, uuid)
}
