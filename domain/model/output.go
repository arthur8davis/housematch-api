package model

import "github.com/google/uuid"

type CreateOutput struct {
	Id *uuid.UUID `json:"id"`
}

type DeleteOutput struct {
	Deleted bool `json:"deleted"`
}

type UpdateOutput struct {
	Updated bool `json:"updated"`
}

type AssignOutput struct {
	Assigned bool `json:"assigned"`
}
