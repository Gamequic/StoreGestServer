package featuresApi

import (
	"storegestserver/pkg/features/auth"
	authservice "storegestserver/pkg/features/auth/service"
	"storegestserver/pkg/features/food"
	foodservice "storegestserver/pkg/features/food/service"
	"storegestserver/pkg/features/money"
	moneyservice "storegestserver/pkg/features/money/service"
	"storegestserver/pkg/features/orders"
	ordersservice "storegestserver/pkg/features/orders/service"
	"storegestserver/pkg/features/photos"
	photoservice "storegestserver/pkg/features/photos/service"
	"storegestserver/pkg/features/users"
	userservice "storegestserver/pkg/features/users/service"

	"github.com/gorilla/mux"
)

func RegisterSubRoutes(router *mux.Router) {
	userservice.InitUsersService()
	authservice.InitAuthService()
	moneyservice.InitMoneyService()
	foodservice.InitFoodService()
	ordersservice.InitOrdersService()
	photoservice.InitPhotosService()

	apiRouter := router.PathPrefix("/api").Subrouter()

	users.RegisterSubRoutes(apiRouter)
	auth.RegisterSubRoutes(apiRouter)
	money.RegisterSubRoutes(apiRouter)
	food.RegisterSubRoutes(apiRouter)
	orders.RegisterSubRoutes(apiRouter)
	photos.RegisterSubRoutes(apiRouter)
}
