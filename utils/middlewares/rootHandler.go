package middlewares

import (
	"net/http"
	"os"
)

func RootHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var correctUsernameRoot = os.Getenv("ROOTUSERNAME")
		var correctPasswordRoot = os.Getenv("ROOTPASSWORD")

		var userSubmittedUsernameRoot = r.Header.Get("RootUsername")
		var userSubmittedPasswordRoot = r.Header.Get("RootPassword")

		//Check if password is correct
		if correctUsernameRoot != userSubmittedUsernameRoot || correctPasswordRoot != userSubmittedPasswordRoot {
			panic(GormError{Code: http.StatusUnauthorized, Message: "Denied root access!", IsGorm: true})
		}

		next.ServeHTTP(w, r)
	})
}
