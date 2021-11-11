package personas

import (
	"fmt"
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
	_, err := s.gtw.GetPersonas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	message := "Se ha registrado correctamente las personas sas."
	urlReturn := fmt.Sprintf("/?message=%s&resource=personas", message)
	http.Redirect(w, r, urlReturn, http.StatusFound)
}
