package model

import "github.com/google/uuid"

type UserRole struct {
	UserID uuid.UUID `json:"user_id"`
	RoleID uuid.UUID `json:"role_id"`
}

type UserRoleOutput struct {
	User User `json:"user"`
	Role Role `json:"role"`
}

type UserRoleOutputs []UserRoleOutput
