package repositories

import (
	"devquest-server/devquest/domain/models"

	"github.com/google/uuid"
)

type UserRepo interface {
	GetUserByID(uuid.UUID) (*models.User, error)
	GetUsersByCompany(uuid.UUID) ([]*models.User, error)
	GetUsersByProject(uuid.UUID) ([]*models.User, error)
	AddUser(*models.User) (*models.User, error)
	UpdateUser(uuid.UUID) error
}