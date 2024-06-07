package entities

import "github.com/google/uuid"

type (
	User struct {
		ID uuid.UUID `json:"id"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Username string `json:"username"`
		PasswordHash string `json:"password_hash"`
		Role Role `json:"role"`
		Company Company `json:"company"`
	}

  Role struct {
		ID uuid.UUID `json:"id"`
		Title string `json:"title"`
	}
)