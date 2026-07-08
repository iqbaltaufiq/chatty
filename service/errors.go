package service

type RoomErrInvalidRoomID struct{}

func (e *RoomErrInvalidRoomID) Error() string {
	return "invalid room id"
}
