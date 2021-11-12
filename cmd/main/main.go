package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/juanmachuca95/migrations_go/api"
)

func main() {

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
