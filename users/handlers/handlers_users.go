package users

import (
	"fmt"
	"net/http"

	gtw "github.com/juanmachuca95/migrations_go/users/gateways"
	users "github.com/juanmachuca95/migrations_go/users/models"
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
	var users []users.User
	users, err := s.gtw.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = s.gtw.CreateUsersSAS(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	message := "Se ha registrado correctamente los usuarios sas."
	urlReturn := fmt.Sprintf("/?message=%s&resource=users", message)
	http.Redirect(w, r, urlReturn, http.StatusFound)
}
