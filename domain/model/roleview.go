package model

import "github.com/google/uuid"

type RoleView struct {
	RoleID       uuid.UUID `json:"role_id"`
	ViewID       uuid.UUID `json:"view_id"`
	ViewOrder    int       `json:"view_order"`
	ViewPosition string    `json:"view_position"`
}

type RoleViewOutput struct {
	Role         Role   `json:"role"`
	View         View   `json:"view"`
	ViewOrder    int    `json:"view_order"`
	ViewPosition string `json:"view_position"`
}

type RoleViewOutputs []RoleViewOutput
