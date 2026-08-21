package handler

import (
	"chatty/entity"
	"chatty/model"
)

type messagePayload struct {
	Message string `json:"message"`
}

type userPayload struct {
	User entity.User `json:"user"`
}

type chatsPayload struct {
	Chats   []entity.Chat `json:"chats"`
	Message string        `json:"message"`
}

type recentChatsPayload struct {
	Chats   []model.RecentChat `json:"chats"`
	Message string             `json:"message"`
}

type friendsPayload struct {
	Friends []model.Friend `json:"friends"`
	Message string         `json:"message"`
}
