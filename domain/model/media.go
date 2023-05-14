package model

import "github.com/google/uuid"

type Media struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	URL  string    `json:"url"`
	Size int64     `json:"size"`
	Type string    `json:"type"`
}

type Medias []Media
