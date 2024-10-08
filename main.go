package main

import (
	"fmt"
	"net/http"
	"os"
	pkg "storegestserver/pkg/database"
	featuresApi "storegestserver/pkg/features"
	"storegestserver/utils"
	"storegestserver/utils/middlewares"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var Logger *zap.Logger

// execute before main
func init() {
	Logger = utils.NewLogger()
	pkg.Logger = Logger
}

func main() {
	defer Logger.Sync() // flushes buffer, if any
	utils.Dotconfig()
	pkg.InitDB()
	mainRouter := mux.NewRouter()
	port := os.Getenv("PORT")

	mainRouter.Use(middlewares.ErrorHandler)
	mainRouter.Use(middlewares.GormErrorHandler)

	// api
	featuresApi.RegisterSubRoutes(mainRouter)

	mainRouter.HandleFunc("/checkhealth", utils.CheckHealth)

	http.Handle("/", mainRouter)
	Logger.Info(fmt.Sprint("Running on 0.0.0.0:", port))
	http.ListenAndServe(fmt.Sprint(":", port), nil)
}
