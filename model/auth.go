package model

type Credential struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
