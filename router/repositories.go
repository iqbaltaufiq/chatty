package router

import (
	"database/sql"

	"chatty/repository"
)

type repositories struct {
	db *sql.DB
}

func NewRepositories(db *sql.DB) *repositories {
	return &repositories{db}
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
	return repository.NewChatRepository(r.db)
}

func (r *repositories) File() repository.FileRepository {
	return repository.NewFileRepository()
}
