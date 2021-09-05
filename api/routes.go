package api

import (
	"github.com/gorilla/mux"
)

func initRoute() *mux.Router {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "<h1>Hola Mundo</h1>")
	})

	return r
}