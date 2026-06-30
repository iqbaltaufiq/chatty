package middleware

import (
	"context"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// Attach chains the passed `middlewares` before executing the `handler` handler.
// The middleware execution order is from left-to-right.
//
// Example: Attach(userHandler, Logging(), Auth())
//
// Order: Logging() -> Auth() -> userHandler
func Attach(handler http.HandlerFunc, middlewares ...Middleware) http.Handler {
	var final http.Handler
	final = http.HandlerFunc(handler)

	midwLength := len(middlewares)
	if midwLength < 1 {
		return final
	}
	for i := midwLength - 1; i >= 0; i-- {
		final = middlewares[i](final)
	}
	return final
}

// FromCtxUsername takes out username from the context.
//
// This getter is necessary since context key is non-exportable.
func FromCtxUsername(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(usernameKey).(string)
	return username, ok
}
