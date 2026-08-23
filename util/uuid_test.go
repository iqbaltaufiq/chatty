package util_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"

	"chatty/util"
)

func TestUUID(test *testing.T) {
	test.Run("generate uuid", func(t *testing.T) {
		newID := util.GenerateUUID()
		if newID == "" {
			fmt.Println("uuid is empty")
			t.Fail()
		}
		fmt.Printf("uuid: %s\n", newID)
	})

	test.Run("convert to binary", func(t *testing.T) {
		const uuidBinaryLength = 16
		newID := util.GenerateUUID()

		binaryUUID, err := util.UUIDtoBinary(newID)
		if err != nil {
			fmt.Println("failed converting string uuid to binary uuid", err.Error())
			t.Fail()
		}
		if len(binaryUUID) != uuidBinaryLength {
			fmt.Println("binary uuid should have length of 16 chars")
			t.Fail()
		}
		fmt.Printf("is length 16? [%v]\n", len(binaryUUID) == uuidBinaryLength)

		_, err = util.UUIDtoBinary("randomtexterrorsrandomtexterrors")
		if err == nil {
			fmt.Println("should throw error when the string is not an uuid")
			t.Fail()
		}
		fmt.Println("this error is expected:", err.Error())
	})

	test.Run("uuid binary to string", func(t *testing.T) {
		newID := util.GenerateUUID()
		binaryUUID, _ := util.UUIDtoBinary(newID)
		stringUUID, err := util.UUIDtoString(binaryUUID)
		if err != nil {
			fmt.Println("error when converting binary uuid to string", err.Error())
		}
		if stringUUID != newID {
			fmt.Println("failed converting binary uuid back to its original string")
			t.Fail()
		}
		fmt.Printf("before: %s\nafter: %s\n", newID, stringUUID)
	})
}

func BenchmarkUUID(bench *testing.B) {
	id := uuid.Must(uuid.NewV7())
	binaryID, _ := id.MarshalBinary()

	bench.Run("generate uuid", func(b *testing.B) {
		for b.Loop() {
			util.GenerateUUID()
		}
	})

	bench.Run("uuid binary to string", func(b *testing.B) {
		for b.Loop() {
			util.UUIDtoString(binaryID)
		}
	})

	bench.Run("uuid string to binary v1", func(b *testing.B) {
		for b.Loop() {
			util.UUIDtoBinary(id.String())
		}
	})
	bench.Run("uuid string to binary v2", func(b *testing.B) {
		for b.Loop() {
			util.UUIDtoBinaryArray(id.String())
		}
	})
}
