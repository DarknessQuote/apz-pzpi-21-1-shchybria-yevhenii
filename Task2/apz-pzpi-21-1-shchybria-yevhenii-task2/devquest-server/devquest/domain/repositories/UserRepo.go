package repositories

import (
	"devquest-server/devquest/domain/models"

	"github.com/google/uuid"
)

type UserRepo interface {
	GetUserByID(uuid.UUID) (models.User, error)
	InsertUser(models.User) error
}