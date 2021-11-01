package personas

import (
	"log"
	"net/http"

	gtw "github.com/juanmachuca95/migrations_go/personas/gateways"
)

type PersonasHTTPServices struct {
	gtw gtw.PersonasGateway
}

func NewPersonasHTTPServices() *PersonasHTTPServices {
	return &PersonasHTTPServices{
		gtw: gtw.NewPersonaGateway(),
	}
}

func (s *PersonasHTTPServices) GetPersonasHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := s.gtw.GetPersonas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println(resp)

	w.WriteHeader(200)
}
