package main 

import (
	
	"hexagonal_go/api"
	"os"
	
)

func main(){
	
	//r := api.InitRoute()

	//http.ListenAndServe(":8080", r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	api.Start(port)

}