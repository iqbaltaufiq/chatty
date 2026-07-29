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
	return repository.NewUserRepository(r.db)
}

func (r *repositories) Room() repository.RoomRepository {
	return repository.NewRoomRepository(r.db)
}

func (r *repositories) Relationship() repository.RelationshipRepository {
	return repository.NewRelationshipRepository(r.db)
}

func (r *repositories) Occupant() repository.OccupantRepository {
	return repository.NewOccupantRepository(r.db)
}

func (r *repositories) Chat() repository.ChatRepository {
	return repository.NewChatRepository(r.db)
}

func (r *repositories) File() repository.FileRepository {
	return repository.NewFileRepository()
}
