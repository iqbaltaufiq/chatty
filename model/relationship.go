package model

type Friend struct {
	UUID     string `json:"uuid,omitempty"`
	Username string `json:"username,omitempty"`
	Fullname string `json:"fullname,omitempty"`
	Photo    string `json:"photo,omitempty"`
	RoomID   string `json:"room_id,omitempty"`
}
