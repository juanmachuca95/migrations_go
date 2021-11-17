package api

import (
	"net/http"

	"github.com/gorilla/mux"

	admins "github.com/juanmachuca95/migrations_go/admins/handlers"
	apoderados "github.com/juanmachuca95/migrations_go/apoderados/handlers"
	escribanos "github.com/juanmachuca95/migrations_go/escribanos/handlers"
	home "github.com/juanmachuca95/migrations_go/home/handlers"
	personas "github.com/juanmachuca95/migrations_go/personas/handlers"
	rentas "github.com/juanmachuca95/migrations_go/rentas/handlers"
	sas "github.com/juanmachuca95/migrations_go/sas/handlers"
	socios "github.com/juanmachuca95/migrations_go/socios/handlers"
	users "github.com/juanmachuca95/migrations_go/users/handlers"
)

func InitRoute() *mux.Router {
	r := mux.NewRouter()

	admins := admins.NewAdminsHTTPService()
	apoderados := apoderados.NewApoderadosHTTPService()
	escribanos := escribanos.NewEscribanosHTTPService()
	home := home.NewHomeHTTPService()
	sas := sas.NewSasHTTPService()
	socios := socios.NewSociosHTTPService()
	rentas := rentas.NewRentasHTTPService()
	users := users.NewUsersHTTPService()
	personas := personas.NewPersonasHTTPServices()

	/*Checkeo servidores*/
	/* r.HandleFunc("/testserver", func(w http.ResponseWriter, r *http.Request) {fmt.Fprint(w, "Hello: "+r.Host)}) */

	// Assets - Estilos - images
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",
		http.FileServer(http.Dir("./assets/"))))

	// AdminHTTServices
	rAdmins := r.PathPrefix("/admins").Subrouter()
	rAdmins.HandleFunc("", admins.GetAdminsHandler).Methods("GET")

	// ApoderadosHTTPServices
	r.HandleFunc("/apoderados", apoderados.GetApoderadosHandler).Methods("GET")

	// EscribanosHTTPService
	r.HandleFunc("/escribanos", escribanos.GetEscribanosHandler).Methods("GET")

	// HomeHTTPServices
	r.HandleFunc("/", home.HomeHandler).Methods("GET")

	// SasHTTPServices
	r.HandleFunc("/sas", sas.GetSasHandler).Methods("GET")

	// SociosHTTPServices
	r.HandleFunc("/socios", socios.GetSociosHandler).Methods("GET")

	// RentasHTTPServices
	r.HandleFunc("/rentas", rentas.GetRentasHandler).Methods("GET")

	// UsersHTTPServices
	rUsers := r.PathPrefix("/users").Subrouter()
	rUsers.HandleFunc("", users.GetUsersHandler).Methods("GET")

	// PersonasHTTPServices
	rPersona := r.PathPrefix("/personas").Subrouter()
	rPersona.HandleFunc("", personas.GetPersonasHandler).Methods("GET")

	return r
}
