package main

type Db struct {
	Users       map[int32]User
	UserIdCount int32
}

type Store interface {
	userExists()
	createUser()
}
