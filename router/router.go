package router

import (
	"geo-location/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// router.HandleFunc("/api/validate", controller.Validate).Methods("POST")
	router.HandleFunc("/api/validate", controller.Validate)

	return router
}
