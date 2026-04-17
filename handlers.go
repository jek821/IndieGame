package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type Server struct {
	store Store
}

func NewServer(store Store) *Server {
	return &Server{store: store}
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req RequestCreateUser
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, Response{OK: false, Message: "failed to read request"})
		return
	}
	if err = json.Unmarshal(body, &req); err != nil {
		writeResponse(w, http.StatusBadRequest, Response{OK: false, Message: "invalid request body"})
		return
	}

	exists, err := s.store.UserExists(req.Username)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, Response{OK: false, Message: "database error"})
		return
	}
	if exists {
		writeResponse(w, http.StatusConflict, Response{OK: false, Code: USERNAME_TAKEN, Message: "username already taken"})
		return
	}

	user, err := s.store.CreateUser(req.Username)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, Response{OK: false, Message: "failed to create user"})
		return
	}

	slog.Info("user created", "username", user.Username, "userId", user.UserId)
	writeResponse(w, http.StatusCreated, Response{OK: true, Code: USER_CREATED, Data: user})
}

func writeResponse(w http.ResponseWriter, status int, res Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
