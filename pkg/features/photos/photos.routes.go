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

func update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	photo := vars["photo"]
	var route string = photoservice.Update(r, photo)
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
	PhotosRouterAuth := PhotosRouter.NewRoute().Subrouter()
	PhotosRouterAuth.Use(middlewares.AuthHandler)

	// Endpoints
	PhotosRouterAuth.HandleFunc("/", create).Methods("POST")
	PhotosRouterAuth.HandleFunc("/{photo}", delete).Methods("DELETE")
	PhotosRouterAuth.HandleFunc("/{photo}", update).Methods("PATCH")

	// Serve static files
	const staticDir string = "/static/"
	PhotosRouter.PathPrefix(staticDir).Handler(http.StripPrefix("/api/photos/static/", http.FileServer(http.Dir("static"))))
}
