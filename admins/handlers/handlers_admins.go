package admins

import (
	"bytes"
	"encoding/json"
	"net/http"

	gtw "github.com/juanmachuca95/migrations_go/admins/gateways"
	models "github.com/juanmachuca95/migrations_go/admins/models"
)

type AdminsHTTPService struct {
	gtw gtw.AdminGateway
}

func NewAdminsHTTPService() *AdminsHTTPService {
	return &AdminsHTTPService{
		gtw: gtw.NewAdminGateway(),
	}
}

func (s *AdminsHTTPService) GetAdminsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var admins []models.Admin
	admins, err := s.gtw.GetAdmins() // Obtiene todas las sas de justicia

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = s.gtw.CreateAdminsSAS(admins)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := map[string]string{"message": "Se ha registrado correctamente los administradores sas."}
	jsonResp, _ := json.Marshal(resp)

	req, _ := http.NewRequest("GET", "/", bytes.NewBuffer(jsonResp))
	req.Header.Set("Content-Type", "application/json")

	http.Redirect(w, req, "/", http.StatusFound)
}
