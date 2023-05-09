package personlocation

import (
	useCasePersonLocation "github.com/arthur8davis/housematch-api/application/usecase/personlocation"
	"github.com/arthur8davis/housematch-api/domain/model"
	storagePersonLocation "github.com/arthur8davis/housematch-api/infrastructure/storage/postgres/personlocation"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCasePersonLocation.New(storagePersonLocation.New(specification.DB))

	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/personsLocations", middlewares...)

	routes.POST("", h.create)
	routes.PUT("/:personID", h.update)
}
