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
	fmt.Println("funciona el handler")
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

/* func main() {
	p1 := NewPersona("juan")
	p2 := NewPersona("larua")

	saludo := p1.Saludar()
	fmt.Println(saludo)

	saludo1 := p2.Saludar()
	fmt.Println(saludo1)

	newMensajesHTTPServices := NewMensajesHTTPServices()
	mensaje := newMensajesHTTPServices.gtw.GetGretting()
	fmt.Println(mensaje)
}
*/
