package model

import "github.com/google/uuid"

type Role struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Order       int       `json:"order"`
}

type Roles []Role
