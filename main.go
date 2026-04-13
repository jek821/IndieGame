package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

var existingUsernames = map[string]struct{}{}
var existingUsers = map[int32]User{}
var userIdCount int32 = 1

type Http struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func newHttp(w http.ResponseWriter, req *http.Request) *Http {
	httpRef := Http{Writer: w, Request: req}
	return &httpRef

}

type RegisterNewUser struct {
	Username string `json:"username"`
}

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

func createResponse(ResponseCode) // not sure the best way to do this

func CreateUser(newUser RegisterNewUser) Response {
	_, exist := existingUsernames[newUser.Username]
	if exist {
		return Response{OK: false, Code: USERNAME_TAKEN, Data: nil, Message: "Please Select New Username This one is taken."}
	}
	newUserStruct := User{
		Username:  newUser.Username,
		UserId:    userIdCount,
		CreatedAt: time.Now().Unix(),
	}
	slog.Info("User created", "username", newUserStruct.Username, "userId", newUserStruct.UserId)
	existingUsers[newUserStruct.UserId] = newUserStruct
	existingUsernames[newUserStruct.Username] = struct{}{}
	userIdCount++
	return Response{OK: true, Code: USER_CREATED, Data: fmt.Sprintf("User ID: %d", newUserStruct.UserId)}
}

func handleCreateUser(w http.ResponseWriter, req *http.Request) {
	var CreateUserReq RegisterNewUser
	requestBody, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	if err = json.Unmarshal(requestBody, &CreateUserReq); err != nil {
		http.Error(w, "Failed to unmarshal request", http.StatusBadRequest)
		return
	}

	response := CreateUser(CreateUserReq)
	// need some way to propogate errors within CreateUser() function to the http.Status
	// Perhaps we could pass the http objects through and instead of returning back here after CreateUser()
	// we just call another function built specifically to write responses back
	responseJson, err := json.Marshal(response)
	_, err = w.Write(responseJson)
	if err != nil {
		http.Error(w, "Failed to Write Response", http.StatusInternalServerError)
		return
	}
}

func sendApiResponse(res Response, httpRef Http) error {
	response, err := json.Marshal(res)
	if err != nil {
		return err
	}
	httpRef.Writer.Write(response)
	return nil

}
func main() {
	initRouteHandlers()
	startServer()
}
