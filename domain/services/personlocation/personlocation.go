package personlocation

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCasePersonLocation interface {
	Create(m model.PersonLocation) (*model.CreateOutput, error)
	Update(personID uuid.UUID, user model.PersonLocation) (*model.UpdateOutput, error)
}
