package entities

import "github.com/google/uuid"

type Achievement struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Points int `json:"points"`
	Project Project `json:"project"`
}