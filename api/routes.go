package api

import (
	"github.com/gorilla/mux"

	login "github.com/juanmachuca95/migrations_go/login/handlers"
	mensajes "github.com/juanmachuca95/migrations_go/mensajes/handlers"

	mdw "github.com/juanmachuca95/migrations_go/internal/middleware"
)

func InitRoute() *mux.Router {

	login := login.NewLoginService()
	mensajes := mensajes.NewMensajesHTTPServices()

	r := mux.NewRouter()
	// MessageHTTPServices
	rMessage := r.PathPrefix("/message").Subrouter()
	rMessage.HandleFunc("", mensajes.GetMensajeHandler).Methods("GET")
	rMessage.Use(mdw.AuthValidToken)

	// LoginHTTPServices
	rLogin := r.PathPrefix("/login").Subrouter()
	rLogin.HandleFunc("", login.LoginHandler).
		Headers("Content-Type", "application/json").
		Methods("POST")

	return r
}
