package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

// Server - server struct
type Server struct {
	// ... TBD...
}

// Run - runs the server
func (s *Server) Run() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:4000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	httpServer := &http.Server{
		Addr:    "localhost:3000",
		Handler: c.Handler(NewRouter(s)),
	}

	fmt.Println("listening on port 3000")
	log.Fatal(httpServer.ListenAndServe())
}

// NewServer - creates a new server with DI
func NewServer() *Server {
	return &Server{}
}
