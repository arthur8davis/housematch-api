package view

import (
	useCaseView "github.com/Melany751/house-match-server/application/usecase/view"
	"github.com/Melany751/house-match-server/domain/model"
	storageView "github.com/Melany751/house-match-server/infrastructure/storage/postgres/view"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCaseView.New(storageView.New(specification.DB))

	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/views", middlewares...)

	routes.GET("/:id", h.getById)
	routes.GET("", h.getAll)
	routes.POST("", h.create)
	routes.PUT("/:id", h.update)
	routes.DELETE("/:id", h.delete)
}
