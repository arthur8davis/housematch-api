package userrole

import (
	useCaseUserRole "github.com/arthur8davis/housematch-api/application/usecase/userrole"
	"github.com/arthur8davis/housematch-api/domain/model"
	storageUserRole "github.com/arthur8davis/housematch-api/infrastructure/storage/postgres/userrole"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCaseUserRole.New(storageUserRole.New(specification.DB))

	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/userRole", middlewares...)

	routes.GET("/user/:userId/role/:roleId", h.getByIds)
	routes.GET("/user/:userId", h.getAll)
	routes.POST("/user/:userId/role/:roleId", h.create)
	routes.DELETE("/user/:userId/role/:roleId", h.delete)
}
