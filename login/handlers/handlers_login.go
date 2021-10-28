package login

import (
	"encoding/json"
	"net/http"

	gtw "github.com/juanmachuca95/migrations_go/login/gateways"
	models "github.com/juanmachuca95/migrations_go/login/models"
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
	var login models.Login

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusPartialContent)
		return
	}

	//fmt.Fprintf(w, "Username: %s - Password: %s", login.Username, login.Password)
	res, err := s.gtw.Login(login.Username, login.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	resp := make(map[string]string)
	resp["token"] = res
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		//log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
	return
}
