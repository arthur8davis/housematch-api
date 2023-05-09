package user

import (
	useCaseUser "github.com/Melany751/house-match-server/application/usecase/user"
	"github.com/Melany751/house-match-server/domain/model"
	storageUser "github.com/Melany751/house-match-server/infrastructure/storage/postgres/user"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCaseUser.New(storageUser.New(specification.DB))

	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/users", middlewares...)

	routes.GET("/:id", h.getById)
	routes.GET("", h.getAll)
	routes.GET("/roles", h.getAllWithRoles)
	routes.POST("", h.create)
	routes.POST("/login", h.login)
	routes.PUT("/:id", h.update)
	routes.DELETE("/:id", h.delete)
}
