package api

import (
	"net/http"

	"github.com/gorilla/mux"

	admins "github.com/juanmachuca95/migrations_go/admins/handlers"
	apoderados "github.com/juanmachuca95/migrations_go/apoderados/handlers"
	home "github.com/juanmachuca95/migrations_go/home/handlers"
	personas "github.com/juanmachuca95/migrations_go/personas/handlers"
	sas "github.com/juanmachuca95/migrations_go/sas/handlers"
	users "github.com/juanmachuca95/migrations_go/users/handlers"
)

func InitRoute() *mux.Router {
	r := mux.NewRouter()

	home := home.NewHomeHTTPService()
	sas := sas.NewSasHTTPService()
	admins := admins.NewAdminsHTTPService()
	apoderados := apoderados.NewApoderadosHTTPService()
	users := users.NewUsersHTTPService()
	personas := personas.NewPersonasHTTPServices()

	/*Checkeo servidores*/
	/* r.HandleFunc("/testserver", func(w http.ResponseWriter, r *http.Request) {fmt.Fprint(w, "Hello: "+r.Host)}) */

	// Assets - Estilos - images
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",
		http.FileServer(http.Dir("./assets/"))))

	// HomeHTTPServices
	r.HandleFunc("/", home.HomeHandler).Methods("GET")

	// SasHTTPServices
	r.HandleFunc("/sas", sas.GetSasHandler).Methods("GET")

	// AdminHTTServices
	rAdmins := r.PathPrefix("/admins").Subrouter()
	rAdmins.HandleFunc("", admins.GetAdminsHandler).Methods("GET")

	r.HandleFunc("/apoderados", apoderados.GetApoderadosHandler).Methods("GET")

	// UsersHTTPServices
	rUsers := r.PathPrefix("/users").Subrouter()
	rUsers.HandleFunc("", users.GetUsersHandler).Methods("GET")

	// PersonasHTTPServices
	rPersona := r.PathPrefix("/personas").Subrouter()
	rPersona.HandleFunc("", personas.GetPersonasHandler).Methods("GET")

	return r
}
