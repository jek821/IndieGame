package main

import (
	"fmt"
	"net/http"
)

func initRouteHandlers() {
	http.HandleFunc("/createUser", handleCreateUser)
}

func startServer() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}
