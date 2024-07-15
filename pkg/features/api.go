package featuresApi

import (
	"storegestserver/pkg/features/auth"
	authservice "storegestserver/pkg/features/auth/service"
	"storegestserver/pkg/features/food"
	foodservice "storegestserver/pkg/features/food/service"
	"storegestserver/pkg/features/money"
	moneyservice "storegestserver/pkg/features/money/service"
	"storegestserver/pkg/features/users"
	userservice "storegestserver/pkg/features/users/service"

	"github.com/gorilla/mux"
)

func RegisterSubRoutes(router *mux.Router) {
	userservice.InitUsersService()
	authservice.InitAuthService()
	moneyservice.InitMoneyService()
	foodservice.InitFoodService()

	apiRouter := router.PathPrefix("/api").Subrouter()

	users.RegisterSubRoutes(apiRouter)
	auth.RegisterSubRoutes(apiRouter)
	money.RegisterSubRoutes(apiRouter)
	food.RegisterSubRoutes(apiRouter)
}
