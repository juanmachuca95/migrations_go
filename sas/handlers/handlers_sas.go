package sas

import (
	"encoding/json"
	"log"
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
	sasforms, err := s.gtw.GetSas()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = s.gtw.CreateSAS(sasforms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(sasforms)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
