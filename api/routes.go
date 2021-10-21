package api

import (
	"github.com/gorilla/mux"

	login "github.com/juanmachuca95/hexagonal_go/login/handlers"
	mensajes "github.com/juanmachuca95/hexagonal_go/mensajes/handlers"
)

func InitRoute() *mux.Router {

	login := login.NewLoginService()
	mensajes := mensajes.NewMensajesHTTPServices()

	r := mux.NewRouter()
	r.HandleFunc("/", mensajes.GetMensajeHandler)

	// LoginHTTPServices
	r.HandleFunc("/login", login.LoginHandler).
		Headers("Content-Type", "application/json").
		Methods("POST")

	return r
}
