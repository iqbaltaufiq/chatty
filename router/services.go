package router

import (
	"database/sql"

	"chatty/service"
)

type services struct {
	db   *sql.DB
	repo *repositories
}

func NewServices(db *sql.DB, repo *repositories) *services {
	return &services{db, repo}
}

// Write down all of the services as method name below.

func (s *services) User() service.UserService {
	return service.NewUserService(s.repo.User())
}

func (s *services) Auth() service.AuthService {
	return service.NewAuthService(s.repo.User(), s.repo.File())
}

func (s *services) Room() service.RoomService {
	return service.NewRoomService(s.repo.Room())
}

func (s *services) Occupant() service.OccupantService {
	return service.NewOccupantService(s.db, s.repo.Occupant())
}

func (s *services) Relationship() service.RelationshipService {
	return service.NewRelationshipService(s.repo.Relationship())
}

func (s *services) Chat() service.ChatService {
	return service.NewChatService(s.repo.Chat(), s.repo.Occupant())
}

func (s *services) File() service.FileService {
	return service.NewFileService(s.repo.File())
}
