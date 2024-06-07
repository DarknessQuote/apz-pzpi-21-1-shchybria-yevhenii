package repositories

import (
	"devquest-server/devquest/domain/entities"

	"github.com/google/uuid"
)

type UserRepo interface {
	GetUserByID(uuid.UUID) (*entities.User, error)
	GetUsersByCompany(uuid.UUID) ([]*entities.User, error)
	GetUsersByProject(uuid.UUID) ([]*entities.User, error)
	AddUser(*entities.User) (*entities.User, error)
	UpdateUser(uuid.UUID) error
}