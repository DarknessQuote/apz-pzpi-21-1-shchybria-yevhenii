package repositories

import (
	"devquest-server/devquest/domain/entities"

	"github.com/google/uuid"
)

type TaskRepo interface {
	GetTasksForProject(uuid.UUID) (*[]entities.Task, error)
	AddTask(entities.Task) (*entities.Task, error)
	UpdateTask(uuid.UUID) error
	DeleteTask(uuid.UUID) error

	AcceptTask(uuid.UUID) error
	CompleteTask(uuid.UUID) error
}