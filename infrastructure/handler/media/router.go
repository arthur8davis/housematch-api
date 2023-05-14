package media

import (
	useCaseMedia "github.com/arthur8davis/housematch-api/application/usecase/media"
	"github.com/arthur8davis/housematch-api/domain/model"
	storageMedia "github.com/arthur8davis/housematch-api/infrastructure/storage/postgres/media"
	"github.com/gin-gonic/gin"
)

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCaseMedia.New(storageMedia.New(specification.DB))

	return newHandler(useCase)
}

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/media", middlewares...)

	routes.POST("/upload", h.upload)
	routes.DELETE("/:id", h.delete)
	routes.GET("/files/*filepath", h.files)
}
