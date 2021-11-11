package sas

import (
	"fmt"
	"net/http"

	models "github.com/juanmachuca95/migrations_go/sas/models"

	gtw "github.com/juanmachuca95/migrations_go/sas/gateways"
)

type SasHTTPService struct {
	gtw gtw.SasGateway
}

func NewSasHTTPService() *SasHTTPService {
	return &SasHTTPService{
		gtw.NewSasGateway(),
	}
}

func (s *SasHTTPService) GetSasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var sasforms []models.Sasform
	sasforms, err := s.gtw.GetSas() // Obtiene todas las sas de justicia

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = s.gtw.CreateSAS(sasforms) // Insert de sas en bd sass
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	message := "Se ha registrado correctamente las sas."
	urlReturn := fmt.Sprintf("/?message=%s&resource=sas", message)
	http.Redirect(w, r, urlReturn, http.StatusFound)
}
