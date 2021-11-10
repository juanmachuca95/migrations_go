package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/juanmachuca95/migrations_go/api"

	page "github.com/juanmachuca95/migrations_go/pages"
)

func main() {

	var message, resource string
	data := page.Page{}
	data.Title = "Juan Machuca: " + message + " - Ha finalizado: " + resource

	initial := "Usuarios"

	data.Steps = []page.Step{
		{Title: "Usuarios", Done: false, Resource: "/users", Key: "users"},
		{Title: "Personas", Done: false, Resource: "/personas", Key: "users"},
		{Title: "Administradores", Done: false, Resource: "/admins", Key: "users"},
		{Title: "Sas", Done: false, Resource: "/sas", Key: "users"},
		{Title: "Socios", Done: false, Resource: "/socios", Key: "users"},
	}

	for index, value := range data.Steps {
		if value.Key == resource {
			value.Done = false
			data.Steps[index+1].Done = true
		}
	}

	log.Println(data.Steps)

	log.Println(initial)

	log.Fatal("terminar")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se han podido cargar las variables de entorno.")
	}

	port := os.Getenv("PORT")
	port2 := os.Getenv("PORT2")

	if port == "" {
		port = "8080"
	}

	if port2 == "" {
		port2 = "8081"
	}

	fmt.Println(port)
	api.Start(port, port2)

}
