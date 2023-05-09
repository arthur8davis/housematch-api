package model

import "github.com/google/uuid"

type Module struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Order       int       `json:"order"`
}

type Modules []Module
