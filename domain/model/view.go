package model

import "github.com/google/uuid"

type View struct {
	ID          uuid.UUID `json:"id"`
	ModuleID    uuid.UUID `json:"module_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Icon        string    `json:"icon"`
}

type Views []View

type ViewOutput struct {
	ID          uuid.UUID `json:"id"`
	Module      Module    `json:"module"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Icon        string    `json:"icon"`
}

type ViewsOutput []ViewOutput
