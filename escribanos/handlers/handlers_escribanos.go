package escribanos

import (
	"encoding/json"
	"net/http"

	gtw "github.com/juanmachuca95/migrations_go/escribanos/gateways"
	models "github.com/juanmachuca95/migrations_go/escribanos/models"
)

type EscribanosHTTPService struct {
	gtw gtw.EscribanoGateway
}

func NewEscribanosHTTPService() *EscribanosHTTPService {
	return &EscribanosHTTPService{
		gtw.NewEscribanoGateway(),
	}
}

func (s *EscribanosHTTPService) GetEscribanosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var escribanosSas []models.EscribanoSAS
	escribanosSas, err := s.gtw.GetEscribanos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = s.gtw.CreateEscribanoSAS(escribanosSas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, _ := json.Marshal(escribanosSas)
	w.Write(jsonResp)

}
