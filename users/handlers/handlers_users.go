package users

import (
	"log"
	"net/http"

	gtw "github.com/juanmachuca95/migrations_go/users/gateways"
)

type UsersHTTPService struct {
	gtw gtw.UsersGateway
}

func NewUsersHTTPService() *UsersHTTPService {
	return &UsersHTTPService{
		gtw: gtw.NewUsersGateway(),
	}
}

func (s *UsersHTTPService) GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := s.gtw.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(resp)

	w.WriteHeader(http.StatusOK)
}
