package main

import "net/http"

func (s *Server) setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/createUser", s.handleCreateUser)
}
