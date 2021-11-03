package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	personas "github.com/juanmachuca95/migrations_go/personas/handlers"
	users "github.com/juanmachuca95/migrations_go/users/handlers"

	/*Template*/
	page "github.com/juanmachuca95/migrations_go/pages"
	template "github.com/juanmachuca95/migrations_go/templates"
)

func InitRoute() *mux.Router {
	r := mux.NewRouter()

	users := users.NewUsersHTTPService()
	personas := personas.NewPersonasHTTPServices()

	/*Checkeo servidores*/
	r.HandleFunc("/testserver", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello: "+r.Host)
	})

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",
		http.FileServer(http.Dir("./assets/"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Template()

		data := page.Page{
			Title: "By Juan Machuca",
			Steps: []page.Step{
				{Title: "Migrar usuarios", Done: true, Resource: "/user"},
				{Title: "Migrar personas", Done: true, Resource: "/personas"},
			},
		}
		tmpl.Execute(w, data)
	})

	// UsersHTTPServices
	rUsers := r.PathPrefix("/users").Subrouter()
	rUsers.HandleFunc("", users.GetUsersHandler).Methods("GET")

	// PersonasHTTPServices
	rPersona := r.PathPrefix("/personas").Subrouter()
	rPersona.HandleFunc("", personas.GetPersonasHandler).Methods("GET")

	return r
}
