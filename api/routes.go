package api

import (
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
		message := r.URL.Query().Get("message")
		resource := r.URL.Query().Get("resource")

		log.Println(message)
		data := page.Page{}
		data.Title = "Juan Machuca: " + message + " - Ha finalizado: " + resource

		initial := "Usuarios"

		data.Steps = []page.Step{
			{Title: "Usuarios", Done: false, Resource: "/users", Key: "users"},
			{Title: "Personas", Done: false, Resource: "/personas", Key: "users"},
			{Title: "Administradores", Done: false, Resource: "/admins", Key: "users"},
			{Title: "Sas", Done: false, Resource: "/sas", Key: "users"},
			{Title: "Socios", Done: false, Resource: "/socios", Key: "users"},
		}

		for index, value := range data.Steps {
			if value.Key == resource {
				value.Done = false
				data.Steps[index+1].Done = true
			}
		}

		log.Println(data.Steps)

		log.Println(initial)

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
