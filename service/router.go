package service

import (
	"github.com/gorilla/mux"
)

// NewRouter gets routes from routes.go and adds them to mux router and returns.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routeList {
		router.HandleFunc(route.path, route.handlerFunc).Methods(route.method).Name(route.name)
	}

	return router
}
