package person

import (
	useCasePerson "github.com/arthur8davis/housematch-api/application/usecase/person"
	"github.com/arthur8davis/housematch-api/domain/model"
	storagePerson "github.com/arthur8davis/housematch-api/infrastructure/storage/postgres/person"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCasePerson.New(storagePerson.New(specification.DB))

	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/persons", middlewares...)

	routes.GET("/:id", h.getById)
	routes.GET("", h.getAll)
	routes.POST("", h.create)
	routes.PUT("/:id", h.update)
	routes.DELETE("/:id", h.delete)
}
