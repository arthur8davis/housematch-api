package model

import "github.com/google/uuid"

type PropertyMedia struct {
	PropertyID uuid.UUID `json:"property_id"`
	MediaID    uuid.UUID `json:"media_id"`
}
