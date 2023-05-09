package transaction

import (
	useCaseTransaction "github.com/arthur8davis/housematch-api/application/usecase/transaction"
	"github.com/arthur8davis/housematch-api/domain/model"
	storageTransaction "github.com/arthur8davis/housematch-api/infrastructure/storage/postgres/transaction"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCaseTransaction.New(storageTransaction.New(specification.DB))

	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/transactions", middlewares...)

	routes.GET("/:id", h.getById)
	//routes.GET("", h.getAll)
	routes.GET("", h.getAllByFilters)
	routes.POST("", h.create)
	routes.PUT("/:id", h.update)
	routes.DELETE("/:id", h.delete)
	routes.GET("/user/:id", h.getByUserId)
}
