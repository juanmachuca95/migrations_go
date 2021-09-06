package web

import (
	"net/http"
	"fmt"
)

func GetGrettingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola desde mi api Golang . Juan</h1>")
}