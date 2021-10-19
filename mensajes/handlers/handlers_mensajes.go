package mensajes

import (
	"fmt"
	"net/http"

	gtw "github.com/juanmachuca95/hexagonal_go/mensajes/gateways"
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

type Interface interface {
	Saludar() string
}

type Persona struct {
	name string
}

func NewPersona(name string) *Persona {
	return &Persona{
		name: name,
	}
}

func (p *Persona) Saludar() string {
	mensaje := fmt.Sprintf("Hola %s Como estás?", p.name)
	return mensaje
}
