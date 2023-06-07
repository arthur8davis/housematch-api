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
	Medias         Medias    `json:"medias"`
}
type PropertiesSecondLevel []PropertySecondLevel

type PropertySecondLevelStorage struct {
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
	Media          Media     `json:"media"`
}

type PropertiesSecondLevelStorage []PropertySecondLevelStorage

func (p PropertiesSecondLevelStorage) PropertiesWithMedias() PropertiesSecondLevel {
	mapTemp := make(map[uuid.UUID]bool)
	mapIndex := make(map[uuid.UUID]int)

	var results PropertiesSecondLevel

	count := 0

	for _, property := range p {
		if _, ok := mapTemp[property.ID]; !ok {
			mapTemp[property.ID] = true
			mapIndex[property.ID] = count
			count++

			var mediasTemp Medias

			results = append(results, PropertySecondLevel{
				ID: property.ID,
				User: User{
					ID:       property.User.ID,
					User:     property.User.User,
					Email:    property.User.Email,
					Theme:    property.User.Theme,
					PersonID: property.User.PersonID,
				},
				Location: Location{
					ID:       property.Location.ID,
					Country:  property.Location.Country,
					City:     property.Location.City,
					Province: property.Location.Province,
					District: property.Location.District,
					Address:  property.Location.Address,
					Lat:      property.Location.Lat,
					Long:     property.Location.Long,
				},
				Description:    property.Description,
				Type:           property.Type,
				Length:         property.Length,
				Width:          property.Width,
				Area:           property.Area,
				Floor:          property.Floor,
				NumberOfFloors: property.NumberOfFloors,
				Rooms:          property.Rooms,
				Bathrooms:      property.Bathrooms,
				Yard:           property.Yard,
				Terrace:        property.Terrace,
				LivingRoom:     property.LivingRoom,
				LaundryRoom:    property.LaundryRoom,
				Kitchen:        property.Kitchen,
				Garage:         property.Garage,
				Medias: append(mediasTemp, Media{
					ID:   property.Media.ID,
					Name: property.Media.Name,
					URL:  property.Media.URL,
					Size: property.Media.Size,
					Type: property.Media.Type,
				}),
			})
		} else {
			results[mapIndex[property.ID]].Medias = append(results[mapIndex[property.ID]].Medias, Media{
				ID:   property.Media.ID,
				Name: property.Media.Name,
				URL:  property.Media.URL,
				Size: property.Media.Size,
				Type: property.Media.Type,
			})
		}
	}

	return results
}

func (pp PropertySecondLevelStorage) PropertyWithMedia() PropertySecondLevel {
	var mediasTemp Medias

	return PropertySecondLevel{
		ID: pp.ID,
		User: User{
			ID:       pp.User.ID,
			User:     pp.User.User,
			Email:    pp.User.Email,
			Theme:    pp.User.Theme,
			PersonID: pp.User.PersonID,
		},
		Location: Location{
			ID:       pp.Location.ID,
			Country:  pp.Location.Country,
			City:     pp.Location.City,
			Province: pp.Location.Province,
			District: pp.Location.District,
			Address:  pp.Location.Address,
			Lat:      pp.Location.Lat,
			Long:     pp.Location.Long,
		},
		Description:    pp.Description,
		Type:           pp.Type,
		Length:         pp.Length,
		Width:          pp.Width,
		Area:           pp.Area,
		Floor:          pp.Floor,
		NumberOfFloors: pp.NumberOfFloors,
		Rooms:          pp.Rooms,
		Bathrooms:      pp.Bathrooms,
		Yard:           pp.Yard,
		Terrace:        pp.Terrace,
		LivingRoom:     pp.LivingRoom,
		LaundryRoom:    pp.LaundryRoom,
		Kitchen:        pp.Kitchen,
		Garage:         pp.Garage,
		Medias: append(mediasTemp, Media{
			ID:   pp.Media.ID,
			Name: pp.Media.Name,
			URL:  pp.Media.URL,
			Size: pp.Media.Size,
			Type: pp.Media.Type,
		}),
	}
}

type PropertyComplete struct {
	Property Property `json:"property"`
	Location Location `json:"location"`
}
