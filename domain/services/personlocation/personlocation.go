package personlocation

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type UseCasePersonLocation interface {
	Create(m model.PersonLocation) (*model.CreateOutput, error)
	Update(personID uuid.UUID, user model.PersonLocation) (*model.UpdateOutput, error)
}
