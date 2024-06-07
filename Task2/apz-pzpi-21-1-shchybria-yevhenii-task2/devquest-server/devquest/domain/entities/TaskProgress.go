package entities

import (
	"time"

	"github.com/google/uuid"
)

type (
	TaskProgress struct {
		ID uuid.UUID `json:"id"`
		Task Task `json:"task"`
		Developer User `json:"developer"`
		Status TaskStatus `json:"status"`
		AcceptedDate time.Time `json:"accepted_time"`
		CompletedDate time.Time `json:"completed_time"`
	}

  TaskStatus struct {
		ID uuid.UUID `json:"id"`
		Name string `json:"status"`
	}
)