package models

import "github.com/google/uuid"

type TaskStatus struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"status"`
}