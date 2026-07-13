package util_test

import (
	"log"
	"testing"

	"chatty/util"
)

func TestHash(test *testing.T) {
	var (
		pwd        = "str0ngpa$$word"
		stringHash string
	)

	test.Run("should return hashed password", func(t *testing.T) {
		h, err := util.HashPassword(pwd)
		if err != nil {
			log.Printf("failed hashing password: %s\n", err.Error())
			t.Fail()
		}
		log.Printf("hash: %s\n", h)
		stringHash = h
	})

	test.Run("should match the hash", func(t *testing.T) {
		if err := util.MatchPassword(pwd, stringHash); err != nil {
			log.Printf("hash does not match: %s\n", err.Error())
			t.Fail()
		}
		log.Println("password match with the hash")
	})
}
