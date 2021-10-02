package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/juanmachuca95/hexagonal_go/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se han podido cargar las variables de entorno.")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println(port)
	api.Start(port)

}
