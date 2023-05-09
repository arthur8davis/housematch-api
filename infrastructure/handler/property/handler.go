package property

import (
	"fmt"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/arthur8davis/housematch-api/domain/services/property"
	"github.com/arthur8davis/housematch-api/domain/services/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

type handler struct {
	useCase property.UseCaseModule
}

func newHandler(useCase property.UseCaseModule) handler {
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

func (h handler) getByUserId(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	m, err := h.useCase.GetByUserId(uid)
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
	var req model.Property
	if err := c.BindJSON(&req); err != nil {
		c.JSON(response.BadRequest(model.ResponseError{Error: fmt.Sprintf("Error read body, error: %s", err.Error())}))
		return
	}

	id, err := h.useCase.Create(req)
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{Error: err.Error()}))
		return
	}

	c.JSON(response.Created(id))
}

func (h handler) createComplete(c *gin.Context) {
	idsMediaString := c.Query("medias_ids")
	var idsMedia []string
	if len(idsMediaString) != 0 {
		idsMedia = strings.Split(idsMediaString, ",")
	}
	var idsMediaUUID []uuid.UUID
	for _, idMedia := range idsMedia {
		uidMedia, err := uuid.Parse(idMedia)
		if err != nil {
			fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
			return
		}
		idsMediaUUID = append(idsMediaUUID, uidMedia)
	}

	var req model.PropertyComplete
	if err := c.BindJSON(&req); err != nil {
		c.JSON(response.BadRequest(model.ResponseError{fmt.Sprintf("Error read body, error: %s", err.Error())}))
		return
	}

	id, err := h.useCase.CreateComplete(req, idsMediaUUID)
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

	var req model.Property
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

func (h handler) updateComplete(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	idsMediaString := c.Query("medias_ids")
	var idsMedia []string
	if len(idsMediaString) != 0 {
		idsMedia = strings.Split(idsMediaString, ",")
	}

	var idsMediaUUID []uuid.UUID
	for _, idMedia := range idsMedia {
		uidMedia, err := uuid.Parse(idMedia)
		if err != nil {
			fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
			return
		}
		idsMediaUUID = append(idsMediaUUID, uidMedia)
	}

	var req model.PropertyComplete
	if err := c.BindJSON(&req); err != nil {
		c.JSON(response.BadRequest(model.ResponseError{fmt.Sprintf("Error read body, error: %s", err.Error())}))
		return
	}

	created, err := h.useCase.UpdateComplete(uid, req, idsMediaUUID)
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
