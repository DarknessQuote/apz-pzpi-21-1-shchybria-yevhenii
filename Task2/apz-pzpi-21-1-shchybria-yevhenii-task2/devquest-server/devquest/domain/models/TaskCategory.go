package models

import "github.com/google/uuid"

type TaskCategory struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}