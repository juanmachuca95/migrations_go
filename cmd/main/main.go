package main 

import (
	
	"github.com/juanmachuca95/hexagonal_go/api"
	"net/http"
	
)

func main(){
	
	r := initRoute()

/* 	srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
 */
	http.ListenAndServe(":8080", r)
	//log.Fatal(http.ListenAndServe(":8081", nil))
}