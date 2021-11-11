package apoderados

import (
	"fmt"
	"net/http"

	gtw "github.com/juanmachuca95/migrations_go/apoderados/gateways"
	models "github.com/juanmachuca95/migrations_go/apoderados/models"
)

type ApoderadosHTTPServices struct {
	gtw gtw.ApoderadoGateway
}

func NewApoderadosHTTPService() *ApoderadosHTTPServices {
	return &ApoderadosHTTPServices{
		gtw: gtw.NewApoderadoGateway(),
	}
}

func (s *ApoderadosHTTPServices) GetApoderadosHandler(w http.ResponseWriter, r *http.Request) {
	var apoderados []models.Apoderado

	apoderados, err := s.gtw.GetApoderados()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = s.gtw.CreateApoderadosSAS(apoderados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	message := "Se ha registrado correctamente los usuarios sas."
	urlReturn := fmt.Sprintf("/?message=%s&resource=apoderados", message)
	http.Redirect(w, r, urlReturn, http.StatusFound)
}
