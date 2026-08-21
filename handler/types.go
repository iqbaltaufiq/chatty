package handler

import "chatty/entity"

type messagePayload struct {
	Message string `json:"message"`
}

type userPayload struct {
	User entity.User `json:"user"`
}
