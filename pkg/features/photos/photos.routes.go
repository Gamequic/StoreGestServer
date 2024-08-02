package photos

import (
	"encoding/json"
	"net/http"
	photoservice "storegestserver/pkg/features/photos/service"
	"storegestserver/utils/middlewares"

	"github.com/gorilla/mux"
)

func create(w http.ResponseWriter, r *http.Request) {
	var route string = photoservice.Create(r)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"route": route,
	})
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var photo string = vars["photo"]
	photoservice.Delete(photo)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"photo": photo,
	})
}

// Register function

func RegisterSubRoutes(router *mux.Router) {
	PhotosRouter := router.PathPrefix("/photos").Subrouter()

	// Middlewares
	PhotosRouter.Use(middlewares.AuthHandler)

	// Endpoints
	PhotosRouter.HandleFunc("/", create).Methods("POST")
	PhotosRouter.HandleFunc("/{photo}", delete).Methods("DELETE")

	// Serve static files
	const staticDir string = "/static/"
	PhotosRouter.PathPrefix(staticDir).Handler(http.StripPrefix("/api/photos/static/", http.FileServer(http.Dir("static"))))
}
