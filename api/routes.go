package api

import (
	"github.com/gorilla/mux"

	mensajes "github.com/juanmachuca95/hexagonal_go/mensajes/web"
)

func InitRoute() *mux.Router {
	mensajes := mensajes.NewMensajesHTTPServices()

	r := mux.NewRouter()
	r.HandleFunc("/", mensajes.GetMensajeHandler)

	return r
}
