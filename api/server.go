package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
)

type server struct {
	*http.Server
}

func newServer(port string, r *mux.Router) *server {
	s := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server{s}
}

// Start runs ListenAndServe on the http.Server with graceful shutdown
func (srv *server) Start() {
	log.Println("starting API cmd")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not listen on %s due to %s", srv.Addr, err.Error())
	}
	log.Printf("cmd is ready to handle requests %s", srv.Addr)
}

