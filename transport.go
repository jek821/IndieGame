package main

import (
	"encoding/json"
	"net/http"
)

type RequestCode int

const (
	CREATE_USER RequestCode = iota
	REQUEST_USER_LOGIN
	REQUEST_TOWNS_VIEW
	REQUEST_TOWN_VIEW
	REQUEST_UPDATE_TOWN
)

type RequestCreateUser struct {
	Username string `json:"username"`
}

type RequestTownsView struct {
	Username   string `json:"username"`
	SessionKey int32  `json:"session_key"`
}

type RequestTownView struct {
	Username   string `json:"username"`
	SessionKey int32  `json:"session_key"`
	TownID     int32  `json:"town_id"`
}

type ResponseCode int

const (
	USER_CREATED ResponseCode = iota
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
