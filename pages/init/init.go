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
		{Title: "Apoderados", Done: false, Resource: "/apoderados", Key: "apoderados"},

		/* {Title: "Rentas", Done: false, Resource: "/rentas", Key: "rentas"}, */
		{Title: "Socios", Done: false, Resource: "/socios", Key: "socios"},
		{Title: "Escribanos", Done: false, Resource: "/escribanos", Key: "escribanos"},

		{Title: "Migración finalizada", Done: false, Resource: "/success", Key: "success"},
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
