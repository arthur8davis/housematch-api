package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Theme    string    `json:"theme"`
	PersonID uuid.UUID `json:"person_id"`
}

type Users []User

type UserSecondLevel struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Theme    string    `json:"theme"`
	Person   Person    `json:"person"`
}

type UsersSecondLevel []UserSecondLevel

type UserWithRolesOutput struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Theme    string    `json:"theme"`
	PersonID uuid.UUID `json:"person_id"`
	Roles    Roles     `json:"roles"`
}

type UsersWithRolesOutput []UserWithRolesOutput

type UserWithRole struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Theme    string    `json:"theme"`
	PersonID uuid.UUID `json:"person_id"`
	Role     Role      `json:"role"`
}

type UsersWithRoles []UserWithRole

func (ur UsersWithRoles) GetUserWithRole() UsersWithRolesOutput {
	var results UsersWithRolesOutput
	mapa := make(map[uuid.UUID]bool)
	mapaIndex := make(map[uuid.UUID]int)
	cont := 0

	for _, element := range ur {
		if _, ok := mapa[element.ID]; !ok {
			mapa[element.ID] = true
			mapaIndex[element.ID] = cont
			cont = cont + 1
			var RolesI Roles
			results = append(results, UserWithRolesOutput{
				ID:       element.ID,
				User:     element.User,
				Password: element.Password,
				Email:    element.Email,
				Theme:    element.Theme,
				PersonID: element.PersonID,
				Roles: append(RolesI, Role{
					ID:          element.Role.ID,
					Name:        element.Role.Name,
					Description: element.Role.Description,
					Order:       element.Role.Order,
				}),
			})
		} else {
			results[mapaIndex[element.ID]].Roles = append(results[mapaIndex[element.ID]].Roles, Role{
				ID:          element.Role.ID,
				Name:        element.Role.Name,
				Description: element.Role.Description,
				Order:       element.Role.Order,
			})
		}
	}

	return results
}

type Login struct {
	User     string `json:"user"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
