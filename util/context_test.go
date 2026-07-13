package util_test

import (
	"context"
	"log"
	"testing"

	"github.com/google/uuid"

	"chatty/util"
)

func TestContext(test *testing.T) {
	ctx := context.Background()

	test.Run("should attach username", func(t *testing.T) {
		username := "user123"
		newCtx := util.SetCtxUsername(ctx, username)
		usernameCtx, ok := util.FromCtxUsername(newCtx)
		if !ok || usernameCtx != username {
			log.Println("failed getting value from context")
			t.Fail()
		}
		log.Printf("username: %s\n", usernameCtx)
	})

	test.Run("should attach uuid", func(t *testing.T) {
		userID := uuid.New().String()
		newCtx := util.SetCtxUUID(ctx, userID)
		userIDCtx, ok := util.FromCtxUUID(newCtx)
		if !ok || userID != userIDCtx {
			log.Println("failed getting value from context")
			t.Fail()
		}
		log.Printf("userID: %s\n", userIDCtx)
	})
}
