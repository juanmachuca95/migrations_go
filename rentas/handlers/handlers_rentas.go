package rentas

import (
	"fmt"
	"net/http"

	gtw "github.com/juanmachuca95/migrations_go/rentas/gateways"
	models "github.com/juanmachuca95/migrations_go/rentas/models"
)

type RentasHTTPService struct {
	gtw gtw.RentasGateway
}

func NewRentasHTTPService() *RentasHTTPService {
	return &RentasHTTPService{
		gtw: gtw.NewRentasGateway(),
	}
}

func (s *RentasHTTPService) GetRentasHandler(w http.ResponseWriter, r *http.Request) {

	var rentas []models.Renta
	rentas, err := s.gtw.GetRentas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = s.gtw.CreateRentasSAS(rentas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	message := "Se ha registrado correctamente los rentas sas."
	urlReturn := fmt.Sprintf("/?message=%s&resource=rentas", message)
	http.Redirect(w, r, urlReturn, http.StatusFound)
}
