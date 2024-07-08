package featuresApi

import (
	"storegestserver/pkg/features/auth"
	authservice "storegestserver/pkg/features/auth/service"
	"storegestserver/pkg/features/users"
	userservice "storegestserver/pkg/features/users/service"

	"github.com/gorilla/mux"
)

func RegisterSubRoutes(router *mux.Router) {
	userservice.InitUsersService()
	authservice.InitAuthService()

	apiRouter := router.PathPrefix("/api").Subrouter()

	users.RegisterSubRoutes(apiRouter)
	auth.RegisterSubRoutes(apiRouter)
}
