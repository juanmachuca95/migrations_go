package home

import (
	"log"
	"net/http"

	/*Template*/
	initpage "github.com/juanmachuca95/migrations_go/pages/init"
	template "github.com/juanmachuca95/migrations_go/templates"
)

type HomeHTTPService struct{}

func NewHomeHTTPService() *HomeHTTPService {
	return &HomeHTTPService{}
}

func (s *HomeHTTPService) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Template()
	message := r.URL.Query().Get("message")
	resource := r.URL.Query().Get("resource")

	log.Println(message, resource)

	initpage := initpage.NewInitPage(resource, message)
	tmpl.Execute(w, initpage)
}
