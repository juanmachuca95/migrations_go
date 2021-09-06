package main 

import (
	
	"hexagonal_go/api"
	"os"
	
)

func main(){
	
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	api.Start(port)

}