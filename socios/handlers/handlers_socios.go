package socios

import (
	"fmt"
	"net/http"

	gtw "github.com/juanmachuca95/migrations_go/socios/gateways"
	models "github.com/juanmachuca95/migrations_go/socios/models"
)

type SociosHTTPServices struct {
	gtw gtw.SociosGateway
}

func NewSociosHTTPService() *SociosHTTPServices {
	return &SociosHTTPServices{
		gtw.NewSociosGateway(),
	}
}

func (s *SociosHTTPServices) GetSociosHandler(w http.ResponseWriter, r *http.Request) {
	var socios []models.Socio
	socios, err := s.gtw.GetSocios()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = s.gtw.CreateSociosSAS(socios)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	message := "Se ha registrado correctamente los socios sas."
	urlReturn := fmt.Sprintf("/?message=%s&resource=socios", message)
	http.Redirect(w, r, urlReturn, http.StatusFound)
}
