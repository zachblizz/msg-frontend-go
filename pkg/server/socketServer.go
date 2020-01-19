package server

import (
	"fmt"
	"net/http"
)

func (s *Server) test(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello worlds")
	return
}
