package util

import "github.com/google/uuid"

// GenerateUUID uses UUID v7 to improve database indexing
// and reduces disk fragmentation.
func GenerateUUID() string {
	newID := uuid.Must(uuid.NewV7())
	return newID.String()
}

func UUIDtoString(b []byte) (string, error) {
	newID, err := uuid.FromBytes(b)
	return newID.String(), err
}

func UUIDtoBinary(s string) ([]byte, error) {
	thisID, err := uuid.Parse(s)
	if err != nil {
		return nil, err
	}
	return thisID.MarshalBinary()
}
