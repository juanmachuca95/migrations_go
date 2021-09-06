package api

import (
	"github.com/gorilla/mux"
	"hexagonal_go/mensajes/web"
)

func InitRoute() *mux.Router {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", web.GetGrettingHandler)

	return r
}