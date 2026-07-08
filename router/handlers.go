package router

import "chatty/handler"

type handlers struct {
	service *services
}

func NewHandlers(service *services) *handlers {
	return &handlers{service}
}

// Write down all of the handlers as method name below.

func (h *handlers) Auth() handler.AuthHandler {
	return handler.NewAuthHandler(
		h.service.Auth(),
		h.service.User(),
	)
}

func (h *handlers) User() handler.UserHandler {
	return handler.NewUserHandler(
		h.service.User(),
	)
}

func (h *handlers) Room() handler.RoomHandler {
	return handler.NewRoomHandler(h.service.Room())
}
