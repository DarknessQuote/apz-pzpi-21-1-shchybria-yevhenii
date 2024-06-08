package repositories

import (
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"

	"github.com/google/uuid"
)

type UserRepo interface {
	GetUserByUsername(username string) (*entities.User, error)
	GetRoleByID(roleID uuid.UUID) (*entities.Role, error)
	CheckUserRole(userID uuid.UUID, roleTitle string) (bool, error)

	InsertUser(user *models.InsertUserDTO) error
}