package repositories

import (
	"devquest-server/devquest/domain/models"

	"github.com/google/uuid"
)

type ProjectRepo interface {
	GetProjectsOfManager(uuid.UUID) ([]*models.Project, error)
	GetProjectsOfDeveloper(uuid.UUID) ([]*models.Project, error)
	AddProject(models.Project) (*models.Project, error)
	UpdateProject(uuid.UUID) error
	DeleteProject(uuid.UUID) error

	AddUserToProject(models.Project, models.User) error
	RemoveUserFromProject(models.Project, models.User) error
	AddAchievementToProject(models.Project, models.Achievement) error
	GiveAchievementToUser(models.Achievement, models.User) error	
}