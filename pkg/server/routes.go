package server

import "net/http"

// Route - the route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// GetRoutes - gets the list of routes for the site
func GetRoutes(s *Server) []Route {
	return []Route{
		Route{"test", "GET", "/test", s.test},
	}
}
