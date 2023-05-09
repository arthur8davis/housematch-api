package module

import (
	useCaseModule "github.com/Melany751/house-match-server/application/usecase/module"
	"github.com/Melany751/house-match-server/domain/model"
	storageModule "github.com/Melany751/house-match-server/infrastructure/storage/postgres/module"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCaseModule.New(storageModule.New(specification.DB))

	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/modules", middlewares...)

	routes.GET("/:id", h.getById)
	routes.GET("", h.getAll)
	routes.POST("", h.create)
	routes.PUT("/:id", h.update)
	routes.DELETE("/:id", h.delete)
}
