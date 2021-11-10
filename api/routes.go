package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	admins "github.com/juanmachuca95/migrations_go/admins/handlers"
	personas "github.com/juanmachuca95/migrations_go/personas/handlers"
	sas "github.com/juanmachuca95/migrations_go/sas/handlers"
	users "github.com/juanmachuca95/migrations_go/users/handlers"

	/*Template*/
	page "github.com/juanmachuca95/migrations_go/pages"
	template "github.com/juanmachuca95/migrations_go/templates"
)

func InitRoute() *mux.Router {
	r := mux.NewRouter()

	sas := sas.NewSasHTTPService()
	admins := admins.NewAdminsHTTPService()
	users := users.NewUsersHTTPService()
	personas := personas.NewPersonasHTTPServices()

	/*Checkeo servidores*/
	/* r.HandleFunc("/testserver", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello: "+r.Host)
	}) */

	// Assets - Estilos - images
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",
		http.FileServer(http.Dir("./assets/"))))

	var home = func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Template()
		var message string
		_ = json.NewDecoder(r.Body).Decode(&message)
		/* if err != nil {
			http.Error(w, err.Error(), 400)
			return
		} */
		log.Println(message)

		data := page.Page{
			Title: "By Juan Machuca",
			Steps: []page.Step{
				{Title: "Usuarios", Done: true, Resource: "/users"},
				{Title: "Personas", Done: false, Resource: "/personas"},
				{Title: "Administradores", Done: false, Resource: "/admins"},
				{Title: "SAS", Done: false, Resource: "/sas"},
			},
		}
		tmpl.Execute(w, data)
	}

	r.HandleFunc("/", home).Methods("GET")

	// SasHTTPServices
	r.HandleFunc("/sas", sas.GetSasHandler).Methods("GET")

	// AdminHTTServices
	rAdmins := r.PathPrefix("/admins").Subrouter()
	rAdmins.HandleFunc("", admins.GetAdminsHandler).Methods("GET")

	// UsersHTTPServices
	rUsers := r.PathPrefix("/users").Subrouter()
	rUsers.HandleFunc("", users.GetUsersHandler).Methods("GET")

	// PersonasHTTPServices
	rPersona := r.PathPrefix("/personas").Subrouter()
	rPersona.HandleFunc("", personas.GetPersonasHandler).Methods("GET")

	return r
}
