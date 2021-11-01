package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	login "github.com/juanmachuca95/migrations_go/login/handlers"
	mensajes "github.com/juanmachuca95/migrations_go/mensajes/handlers"
	personas "github.com/juanmachuca95/migrations_go/personas/handlers"
	users "github.com/juanmachuca95/migrations_go/users/handlers"

	mdw "github.com/juanmachuca95/migrations_go/internal/middleware"
)

func InitRoute() *mux.Router {
	r := mux.NewRouter()

	login := login.NewLoginService()
	mensajes := mensajes.NewMensajesHTTPServices()
	users := users.NewUsersHTTPService()
	personas := personas.NewPersonasHTTPServices()

	/*Checkeo servidores*/
	r.HandleFunc("/testserver", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello: "+r.Host)
	})

	// MessageHTTPServices
	rMessage := r.PathPrefix("/message").Subrouter()
	rMessage.HandleFunc("", mensajes.GetMensajeHandler).Methods("GET")
	rMessage.Use(mdw.AuthValidToken)

	// UsersHTTPServices
	rUsers := r.PathPrefix("/users").Subrouter()
	rUsers.HandleFunc("", users.GetUsersHandler).Methods("GET")

	// PersonasHTTPServices
	rPersona := r.PathPrefix("/personas").Subrouter()
	rPersona.HandleFunc("", personas.GetPersonasHandler).Methods("GET")

	// LoginHTTPServices
	rLogin := r.PathPrefix("/login").Subrouter()
	rLogin.HandleFunc("", login.LoginHandler).
		Headers("Content-Type", "application/json").
		Methods("POST")

	return r
}
