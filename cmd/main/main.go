package main 

import (
	
	api "hexagonal_go/api"
	"net/http"
	
)

func main(){
	
	r := api.InitRoute()
	http.ListenAndServe(":8080", r)

}