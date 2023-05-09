package person

import (
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/Melany751/house-match-server/domain/services/person"
	"github.com/Melany751/house-match-server/domain/services/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	useCase person.UseCasePerson
}

func newHandler(useCase person.UseCasePerson) handler {
	return handler{useCase}
}

func (h handler) getById(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	m, err := h.useCase.GetById(uid)
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{Error: err.Error()}))
		return
	}
	c.JSON(response.OK(m))
}

func (h handler) getAll(c *gin.Context) {
	ms, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{Error: err.Error()}))
		return
	}

	c.JSON(response.OK(ms))
}

func (h handler) create(c *gin.Context) {
	var req model.Person
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
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	var req model.Person
	if err := c.BindJSON(&req); err != nil {
		c.JSON(response.BadRequest(model.ResponseError{Error: fmt.Sprintf("Error read body, error: %s", err.Error())}))
		return
	}

	created, err := h.useCase.Update(uid, req)
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{Error: err.Error()}))
		return
	}

	c.JSON(response.Updated(created))
}

func (h handler) delete(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	deleted, err := h.useCase.Delete(uid)
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{Error: err.Error()}))
		return
	}

	c.JSON(response.Deleted(deleted))
}
