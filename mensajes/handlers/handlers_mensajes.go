package mensajes

import (
	"fmt"
	"net/http"

	gtw "github.com/juanmachuca95/migrations_go/mensajes/gateways"
)

type MensajesHTTPServices struct {
	gtw gtw.MensajesGateway
}

func NewMensajesHTTPServices() *MensajesHTTPServices {
	return &MensajesHTTPServices{
		gtw.NewMensajesGateway(),
	}
}

func (s *MensajesHTTPServices) GetMensajeHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	var resp string
	resp = s.gtw.GetGretting()

	fmt.Fprintf(w, "%s\n", resp)
}
