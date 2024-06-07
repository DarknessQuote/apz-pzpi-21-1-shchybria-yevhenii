package repositories

import (
	"devquest-server/devquest/domain/entities"

	"github.com/google/uuid"
)

type ProjectRepo interface {
	GetProjectsOfManager(uuid.UUID) ([]*entities.Project, error)
	GetProjectsOfDeveloper(uuid.UUID) ([]*entities.Project, error)
	AddProject(entities.Project) (*entities.Project, error)
	UpdateProject(uuid.UUID) error
	DeleteProject(uuid.UUID) error

	AddUserToProject(entities.Project, entities.User) error
	RemoveUserFromProject(entities.Project, entities.User) error
	AddAchievementToProject(entities.Project, entities.Achievement) error
	GiveAchievementToUser(entities.Achievement, entities.User) error	
}