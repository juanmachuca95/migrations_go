package main 

import (
	"log"
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "<h1>Hola Mundo</h1>")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}