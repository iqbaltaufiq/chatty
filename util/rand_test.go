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

func BenchmarkRand(bench *testing.B) {
	bench.Run("generate 6 digits alphanumeric", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			util.GenerateAlphanumeric(6)
		}
	})
}
