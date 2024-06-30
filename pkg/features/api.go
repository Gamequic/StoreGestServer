package featuresApi

import (
	"storegestserver/pkg/features/users"

	"github.com/gorilla/mux"
)

func RegisterSubRoutes(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()

	users.RegisterSubRoutes(apiRouter)

}
