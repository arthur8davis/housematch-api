package property

import (
	useCaseProperty "github.com/arthur8davis/housematch-api/application/usecase/property"
	"github.com/arthur8davis/housematch-api/domain/model"
	storageProperty "github.com/arthur8davis/housematch-api/infrastructure/storage/postgres/property"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCaseProperty.New(storageProperty.New(specification.DB))

	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/properties", middlewares...)

	routes.GET("/:id", h.getById)
	routes.GET("", h.getAll)
	routes.POST("", h.create)
	routes.POST("/create", h.createComplete)
	routes.PUT("/:id", h.update)
	routes.PUT("/update/:id", h.updateComplete)
	routes.DELETE("/:id", h.delete)
	routes.GET("/user/:id", h.getByUserId)
}
