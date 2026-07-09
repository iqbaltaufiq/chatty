package router

import "chatty/service"

type services struct {
	repo *repositories
}

func NewServices(repo *repositories) *services {
	return &services{repo}
}

// Write down all of the services as method name below.

func (s *services) User() service.UserService {
	return service.NewUserService(s.repo.User())
}

func (s *services) Auth() service.AuthService {
	return service.NewAuthService(s.repo.User())
}

func (s *services) Room() service.RoomService {
	return service.NewRoomService(s.repo.Room())
}

func (s *services) Occupant() service.OccupantService {
	return service.NewOccupantService(s.repo.Occupant())
}

func (s *services) Relationship() service.RelationshipService {
	return service.NewRelationshipService(s.repo.Relationship())
}
