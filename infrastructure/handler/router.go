package handler

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	routerLocationPerson "github.com/arthur8davis/housematch-api/infrastructure/handler/location"
	routerMedia "github.com/arthur8davis/housematch-api/infrastructure/handler/media"
	routerModule "github.com/arthur8davis/housematch-api/infrastructure/handler/module"
	routerPerson "github.com/arthur8davis/housematch-api/infrastructure/handler/person"
	routerPersonLocation "github.com/arthur8davis/housematch-api/infrastructure/handler/personlocation"
	routerProperty "github.com/arthur8davis/housematch-api/infrastructure/handler/property"
	routerRole "github.com/arthur8davis/housematch-api/infrastructure/handler/role"
	routerRoleView "github.com/arthur8davis/housematch-api/infrastructure/handler/roleview"
	routerTransaction "github.com/arthur8davis/housematch-api/infrastructure/handler/transaction"
	routerUser "github.com/arthur8davis/housematch-api/infrastructure/handler/user"
	routerUserRole "github.com/arthur8davis/housematch-api/infrastructure/handler/userrole"
	routerView "github.com/arthur8davis/housematch-api/infrastructure/handler/view"
)

func InitRoutes(specification model.RouterSpecification) {
	// User
	routerUser.NewRouter(specification)
	// Person
	routerPerson.NewRouter(specification)
	// Location
	routerLocationPerson.NewRouter(specification)
	// PersonLocation
	routerPersonLocation.NewRouter(specification)
	// Role
	routerRole.NewRouter(specification)
	// Module
	routerModule.NewRouter(specification)
	// View
	routerView.NewRouter(specification)
	// RoleView
	routerRoleView.NewRouter(specification)
	// UserRole
	routerUserRole.NewRouter(specification)
	// Property
	routerProperty.NewRouter(specification)
	// Transaction
	routerTransaction.NewRouter(specification)
	// Media
	routerMedia.NewRouter(specification)
}
