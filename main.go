package main

import (
	"fmt"
	pkg "storegestserver/pkg/database"
	"storegestserver/utils"

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

	fmt.Println(Logger)
	utils.Dotconfig()
	pkg.InitDB()
}
