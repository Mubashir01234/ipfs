package routes

import (
	"ipfs/controllers"

	"github.com/gorilla/mux"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()
	// ipfs endpoint
	file := api.PathPrefix("/file").Subrouter()
	file.HandleFunc("/upload", controllers.UploadFile).Methods("POST")

	return router
}
