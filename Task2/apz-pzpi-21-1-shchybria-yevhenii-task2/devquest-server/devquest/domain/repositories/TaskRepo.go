package repositories

import (
	"devquest-server/devquest/domain/models"

	"github.com/google/uuid"
)

type TaskRepo interface {
	GetTasksForProject(uuid.UUID) (*[]models.Task, error)
	AddTask(models.Task) (*models.Task, error)
	UpdateTask(uuid.UUID) error
	DeleteTask(uuid.UUID) error

	AcceptTask(uuid.UUID) error
	CompleteTask(uuid.UUID) error
}