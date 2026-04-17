package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Code RequestCode     `json:"code"`
	Data json.RawMessage `json:"data"`
}

type Response struct {
	OK      bool         `json:"ok"`
	Code    ResponseCode `json:"code"`
	Data    any          `json:"data"`
	Message string       `json:"message"`
}

func main() {
	store, err := NewSqliteStore("game.db")
	if err != nil {
		log.Fatal("failed to open database:", err)
	}

	server := NewServer(store)

	mux := http.NewServeMux()
	server.setupRoutes(mux)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
