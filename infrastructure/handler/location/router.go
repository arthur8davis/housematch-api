package location

import (
	useCaseLocationPerson "github.com/Melany751/house-match-server/application/usecase/location"
	"github.com/Melany751/house-match-server/domain/model"
	storageLocationPerson "github.com/Melany751/house-match-server/infrastructure/storage/postgres/locationperson"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCaseLocationPerson.New(storageLocationPerson.New(specification.DB))

	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/locations", middlewares...)

	routes.GET("/:id", h.getById)
	routes.GET("", h.getAll)
	routes.POST("", h.create)
	routes.PUT("/:id", h.update)
	routes.DELETE("/:id", h.delete)
}
