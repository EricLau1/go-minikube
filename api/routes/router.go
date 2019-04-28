package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	for _, route := range Routes {
		r.HandleFunc(route.Pattern, route.Handler).Methods(route.Method)
	}
	return r
}