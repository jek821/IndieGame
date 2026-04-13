package main

type User struct {
	Username  string `json:"username"`
	UserId    int32  `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
}
