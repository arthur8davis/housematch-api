package model

import (
	"github.com/google/uuid"
)

type Property struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
	LocationID     uuid.UUID `json:"location_id"`
	Description    string    `json:"description"`
	Type           string    `json:"type"`
	Length         float64   `json:"length"`
	Width          float64   `json:"width"`
	Area           float64   `json:"area"`
	Floor          float64   `json:"floor"`
	NumberOfFloors float64   `json:"number_of_floors"`
	Rooms          int       `json:"rooms"`
	Bathrooms      int       `json:"bathrooms"`
	Yard           int       `json:"yard"`
	Terrace        int       `json:"terrace"`
	LivingRoom     int       `json:"living_room"`
	LaundryRoom    int       `json:"laundry_room"`
	Kitchen        int       `json:"kitchen"`
	Garage         int       `json:"garage"`
}

type Properties []Property

type PropertySecondLevel struct {
	ID             uuid.UUID `json:"id"`
	User           User      `json:"user"`
	Location       Location  `json:"location"`
	Description    string    `json:"description"`
	Type           string    `json:"type"`
	Length         float64   `json:"length"`
	Width          float64   `json:"width"`
	Area           float64   `json:"area"`
	Floor          float64   `json:"floor"`
	NumberOfFloors float64   `json:"number_of_floors"`
	Rooms          int       `json:"rooms"`
	Bathrooms      int       `json:"bathrooms"`
	Yard           int       `json:"yard"`
	Terrace        int       `json:"terrace"`
	LivingRoom     int       `json:"living_room"`
	LaundryRoom    int       `json:"laundry_room"`
	Kitchen        int       `json:"kitchen"`
	Garage         int       `json:"garage"`
}

type PropertiesSecondLevel []PropertySecondLevel

type PropertySecondLevelWithoutUser struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
	Location       Location  `json:"location"`
	Description    string    `json:"description"`
	Type           string    `json:"type"`
	Length         float64   `json:"length"`
	Width          float64   `json:"width"`
	Area           float64   `json:"area"`
	Floor          float64   `json:"floor"`
	NumberOfFloors float64   `json:"number_of_floors"`
	Rooms          int       `json:"rooms"`
	Bathrooms      int       `json:"bathrooms"`
	Yard           int       `json:"yard"`
	Terrace        int       `json:"terrace"`
	LivingRoom     int       `json:"living_room"`
	LaundryRoom    int       `json:"laundry_room"`
	Kitchen        int       `json:"kitchen"`
	Garage         int       `json:"garage"`
}

type PropertiesSecondLevelWithoutUser []PropertySecondLevelWithoutUser

type PropertyComplete struct {
	Property Property `json:"property"`
	Location Location `json:"location"`
}
