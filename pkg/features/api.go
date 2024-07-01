package featuresApi

import (
	"storegestserver/pkg/features/users"
	userservice "storegestserver/pkg/features/users/service"

	"github.com/gorilla/mux"
)

func RegisterSubRoutes(router *mux.Router) {
	userservice.InitUsersService()

	apiRouter := router.PathPrefix("/api").Subrouter()

	users.RegisterSubRoutes(apiRouter)

}
