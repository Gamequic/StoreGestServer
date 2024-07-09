package middlewares

import (
	"net/http"
	"storegestserver/utils"

	"go.uber.org/zap"
)

var logger *zap.Logger

type AppHandler func(http.ResponseWriter, *http.Request) error

type GormError struct {
	Code    int
	Message string
	IsGorm  bool
}

func GormErrorHandler(next http.Handler) http.Handler {
	logger = utils.NewLogger()
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if r, ok := err.(GormError); ok { // It is a gorm error
					var Message string = "Controled panic occurred from " + request.RemoteAddr + " to " + request.URL.Path
					logger.Error(Message, zap.Any("error", err))
					http.Error(w, r.Message, r.Code)
				} else {
					panic(err)
				}
			}
		}()

		next.ServeHTTP(w, request)
	})
}

func ErrorHandler(next http.Handler) http.Handler {
	logger = utils.NewLogger()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic occurred", zap.Any("error", err))
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
