package main

import (
	"os"

	"github.com/juanmachuca95/hexagonal_go/api"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	api.Start(port)

}
