package roleview

import (
	useCaseRoleView "github.com/Melany751/house-match-server/application/usecase/roleview"
	"github.com/Melany751/house-match-server/domain/model"
	storageRoleView "github.com/Melany751/house-match-server/infrastructure/storage/postgres/roleview"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCaseRoleView.New(storageRoleView.New(specification.DB))

	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/roleView", middlewares...)

	routes.GET("/role/:roleID/view/:viewID", h.getByIds)
	routes.GET("", h.getAll)
	routes.POST("", h.assignment)
	routes.PUT("", h.update)
	routes.DELETE("/role/:roleID/view/:viewID", h.delete)
}
