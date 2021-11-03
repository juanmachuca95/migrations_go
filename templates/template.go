package templates

import (
	"html/template"
	"log"
)

func Template() *template.Template {

	// Template
	tpl, err := template.ParseFiles(
		"../../templates/app/index.html",
		"../../templates/parts/header.html")

	if err != nil {
		log.Fatalf("Error al complicar archivos : %s", err)
	}

	return tpl
}
