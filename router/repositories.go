package router

import (
	"chatty/repository"
)

type repositories struct{}

func NewRepositories() *repositories {
	return &repositories{}
}

// Write down each of the repository as a method name below.

func (r *repositories) User() repository.UserRepository {
	return repository.NewUserRepository()
}

func (r *repositories) Room() repository.RoomRepository {
	return repository.NewRoomRepository()
}

func (r *repositories) Relationship() repository.RelationshipRepository {
	return repository.NewRelationshipRepository()
}

func (r *repositories) Occupant() repository.OccupantRepository {
	return repository.NewOccupantRepository()
}

func (r *repositories) Chat() repository.ChatRepository {
	return repository.NewChatRepository()
}

func (r *repositories) File() repository.FileRepository {
	return repository.NewFileRepository()
}
