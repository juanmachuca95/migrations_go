package pages

import (
	"log"

	models "github.com/juanmachuca95/migrations_go/pages/models"
)

func NewInitPage(resource, message string) models.Page {

	var selected string
	var migrationInit = "users"
	var title = "Migración - CODEX.SA - Dev. Machuca Juan"
	data := models.Page{}

	data.Title = title
	data.Message = message

	/* SET DE MIGRACIONES POR PASOS ORDENADOS */
	data.Steps = []models.Step{
		{Title: "Usuarios", Done: false, Resource: "/users", Key: "users"},
		{Title: "Personas", Done: false, Resource: "/personas", Key: "personas"},
		{Title: "Sas", Done: false, Resource: "/sas", Key: "sas"},
		{Title: "Administradores", Done: false, Resource: "/admins", Key: "admins"},
		{Title: "Socios", Done: false, Resource: "/socios", Key: "socios"},
	}

	if resource != "" {
		for index, value := range data.Steps {
			if value.Key == resource {
				value.Done = false
				data.Steps[index+1].Done = true
			}
		}

		for _, value := range data.Steps {
			if value.Done {
				selected = value.Resource
			}
		}
	} else {
		/* Establecemos un inicio de migración users al principio */
		for index, value := range data.Steps {
			if value.Key == migrationInit {
				data.Steps[index].Done = true
				selected = value.Resource
			}
		}

	}

	log.Printf("Recurso elegido es: %s", selected)

	return data
}