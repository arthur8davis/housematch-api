package transaction

import (
	"fmt"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/arthur8davis/housematch-api/domain/services/response"
	"github.com/arthur8davis/housematch-api/domain/services/transaction"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	useCase transaction.UseCaseTransaction
}

func newHandler(useCase transaction.UseCaseTransaction) handler {
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

func (h handler) getAllByFilters(c *gin.Context) {
	typeProperty := c.Query("typeProperty")
	rooms := c.Query("rooms")
	bathrooms := c.Query("bathrooms")
	minArea := c.Query("minArea")
	maxArea := c.Query("maxArea")
	typeTransaction := c.Query("typeTransaction")
	maxCost := c.Query("maxCost")
	minCost := c.Query("minCost")
	country := c.Query("country")
	province := c.Query("province")
	district := c.Query("district")

	fieldsForValidate := map[string]string{
		"typeProperty":    typeProperty,
		"rooms":           rooms,
		"bathrooms":       bathrooms,
		"maxArea":         maxArea,
		"minArea":         minArea,
		"typeTransaction": typeTransaction,
		"maxCost":         maxCost,
		"minCost":         minCost,
		"country":         country,
		"province":        province,
		"district":        district,
	}

	ms, err := h.useCase.GetAllByFilters(fieldsForValidate)
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{Error: err.Error()}))
		return
	}
	c.JSON(response.OK(ms))
}

func (h handler) create(c *gin.Context) {
	var req model.Transaction
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

func (h handler) update(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	var req model.Transaction
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
