package login

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	gtw "github.com/juanmachuca95/hexagonal_go/login/gateways"
)

type LoginHTTPServices struct {
	gtw gtw.LoginGateways
}

func NewLoginService() *LoginHTTPServices {
	return &LoginHTTPServices{
		gtw.NewLoginGateways(),
	}
}

func (s *LoginHTTPServices) LoginHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]
	password := vars["password"]

	res, err := s.gtw.Login(username, password)
	if err != nil {
		fmt.Fprintf(w, "Authentication failed")
		return
	}

	fmt.Fprintf(w, "Authentication Success - Token: %s", res)
	return
}
