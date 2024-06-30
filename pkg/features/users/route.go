package users

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var UsersRouter *mux.Router

func main() {
	UsersRouter = mux.NewRouter()
	UsersRouter.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) { fmt.Println("Hello") }).Methods("GET")
	http.ListenAndServe(":8080", r)
}
