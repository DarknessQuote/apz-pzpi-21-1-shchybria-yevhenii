package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Project Project `json:"project"`
	Category TaskCategory `json:"category"`
	ExpectedTime time.Time `json:"expected_time"`
}