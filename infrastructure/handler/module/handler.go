package module

import (
	"fmt"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/arthur8davis/housematch-api/domain/services/module"
	"github.com/arthur8davis/housematch-api/domain/services/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	useCase module.UseCaseModule
}

func newHandler(useCase module.UseCaseModule) handler {
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
		c.JSON(response.Wrong(model.ResponseError{err.Error()}))
		return
	}
	c.JSON(response.OK(m))
}

func (h handler) getAll(c *gin.Context) {
	ms, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{err.Error()}))
		return
	}
	c.JSON(response.OK(ms))
}

func (h handler) create(c *gin.Context) {
	var req model.Module
	if err := c.BindJSON(&req); err != nil {
		c.JSON(response.BadRequest(model.ResponseError{fmt.Sprintf("Error read body, error: %s", err.Error())}))
		return
	}

	id, err := h.useCase.Create(req)
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{err.Error()}))
		return
	}

	c.JSON(response.Created(id))
}

func (h handler) update(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	var req model.Module
	if err := c.BindJSON(&req); err != nil {
		c.JSON(response.BadRequest(model.ResponseError{fmt.Sprintf("Error read body, error: %s", err.Error())}))
		return
	}

	created, err := h.useCase.Update(uid, req)
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{err.Error()}))
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
		c.JSON(response.Wrong(model.ResponseError{err.Error()}))
		return
	}

	c.JSON(response.Deleted(deleted))
}
