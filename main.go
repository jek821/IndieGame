package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

var existingUsernames = map[string]struct{}{}
var existingUsers = map[int32]User{}
var userIdCount int32 = 1

type RegisterNewUser struct {
	Username string `json:"username"`
}

type User struct {
	username  string
	userId    int32
	createdAt int64
}

type APIResponse struct {
	OK   bool   `json:"ok"`
	Code string `json:"code"`
	Data string `json:"data"`
}

func CreateUser(newUser RegisterNewUser) APIResponse {
	_, exist := existingUsernames[newUser.Username]
	if exist {
		return APIResponse{OK: false, Code: "USERNAME_TAKEN", Data: "Please Select New Username This one is taken."}
	}
	newUserStruct := User{
		username:  newUser.Username,
		userId:    userIdCount,
		createdAt: time.Now().Unix(),
	}
	existingUsers[newUserStruct.userId] = newUserStruct
	existingUsernames[newUserStruct.username] = struct{}{}
	userIdCount++
	return APIResponse{OK: true, Code: "USER_CREATED", Data: fmt.Sprintf("User ID: %d", newUserStruct.userId)}
}

func handleCreateUser(w http.ResponseWriter, req *http.Request) {
	var CreateUserReq RegisterNewUser
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	if err = json.Unmarshal(data, &CreateUserReq); err != nil {
		http.Error(w, "Failed to unmarshal request", http.StatusInternalServerError)
		return
	}

	response := CreateUser(CreateUserReq)
	responseJson, err := json.Marshal(response)
	_, err = w.Write(responseJson)
	if err != nil {
		http.Error(w, "Failed to Write Response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/createUser", handleCreateUser)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
