package personlocation

import (
	"fmt"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/arthur8davis/housematch-api/domain/services/personlocation"
	"github.com/arthur8davis/housematch-api/domain/services/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	useCase personlocation.UseCasePersonLocation
}

func newHandler(useCase personlocation.UseCasePersonLocation) handler {
	return handler{useCase}
}

func (h handler) create(c *gin.Context) {
	var req model.PersonLocation
	if err := c.BindJSON(&req); err != nil {
		c.JSON(response.BadRequest(model.ResponseError{Error: fmt.Sprintf("Error read body, error: %s", err.Error())}))
		return
	}

	m, err := h.useCase.Create(req)
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{Error: err.Error()}))
		return
	}

	c.JSON(response.Created(m))
}

func (h handler) update(c *gin.Context) {
	personID := c.Param("personID")
	personUid, err := uuid.Parse(personID)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	var req model.PersonLocation
	if err := c.BindJSON(&req); err != nil {
		c.JSON(response.BadRequest(model.ResponseError{Error: fmt.Sprintf("Error read body, error: %s", err.Error())}))
		return
	}

	created, err := h.useCase.Update(personUid, req)
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{Error: err.Error()}))
		return
	}

	c.JSON(response.Updated(created))
}
