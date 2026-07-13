package util_test

import (
	"fmt"
	"testing"

	"chatty/util"
)

func TestRand(test *testing.T) {
	const roomIDLength = 6
	randomAlphanum := util.GenerateAlphanumeric(roomIDLength)

	test.Run("should have length of 6", func(t *testing.T) {
		length := len(randomAlphanum)
		if length != 6 {
			fmt.Printf("expected length: %d, actual: %d\n", roomIDLength, length)
			t.Fail()
		}
		fmt.Printf("roomID: %s\n", randomAlphanum)
		fmt.Printf("length: %d\n", length)
	})
}
