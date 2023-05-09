package role

import (
	useCaseRole "github.com/Melany751/house-match-server/application/usecase/role"
	"github.com/Melany751/house-match-server/domain/model"
	storageRole "github.com/Melany751/house-match-server/infrastructure/storage/postgres/role"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCaseRole.New(storageRole.New(specification.DB))

	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/roles", middlewares...)

	routes.GET("/:id", h.getById)
	routes.GET("", h.getAll)
	routes.POST("", h.create)
	routes.PUT("/:id", h.update)
	routes.DELETE("/:id", h.delete)
}
