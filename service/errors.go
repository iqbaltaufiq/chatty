package service

type RoomErrInvalidRoomID struct{}

func (e *RoomErrInvalidRoomID) Error() string {
	return "invalid room id"
}

type RowsErrAlreadyExists struct{}

func (e *RowsErrAlreadyExists) Error() string {
	return "expected to be absent, but resource already exists in database"
}
