package main

import (
	"encoding/json"
	"net/http"
)

type RequestCode int

const (
	CREATE_USER = iota
)

type RequestCreateUser struct {
	Username string `json:"username"`
}

type ResponseCode int

const (
	USER_CREATED = iota
	USERNAME_TAKEN
)

func unpackRequest(req *http.Request) (any, error) {
	var body Request
	json.NewDecoder(req.Body).Decode(&body)

	switch body.Code {
	case CREATE_USER:
		var request RequestCreateUser
		json.Unmarshal(body.Data, &request)
		return request, nil

	}
	return nil, nil
}
