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

func UUIDtoBinaryArray(s string) ([16]byte, error) {
	var binaryID [16]byte
	thisID, err := uuid.Parse(s)
	if err != nil {
		return binaryID, err
	}
	sl, err := thisID.MarshalBinary()
	if err != nil {
		return binaryID, err
	}
	for i := range binaryID {
		binaryID[i] = sl[i]
	}
	return binaryID, err
}
