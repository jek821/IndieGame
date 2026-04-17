package main

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

type Store interface {
	UserExists(username string) (bool, error)
	CreateUser(username string) (User, error)
	GetUserByUsername(username string) (User, error)
}

type SqliteStore struct {
	db *sql.DB
}

func NewSqliteStore(path string) (*SqliteStore, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	store := &SqliteStore{db: db}
	if err := store.migrate(); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *SqliteStore) migrate() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		username   TEXT    UNIQUE NOT NULL,
		created_at INTEGER NOT NULL
	)`)
	return err
}

func (s *SqliteStore) UserExists(username string) (bool, error) {
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	return count > 0, err
}

func (s *SqliteStore) CreateUser(username string) (User, error) {
	now := time.Now().Unix()
	result, err := s.db.Exec("INSERT INTO users (username, created_at) VALUES (?, ?)", username, now)
	if err != nil {
		return User{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}
	return User{UserId: int32(id), Username: username, CreatedAt: now}, nil
}

func (s *SqliteStore) GetUserByUsername(username string) (User, error) {
	var u User
	err := s.db.QueryRow("SELECT id, username, created_at FROM users WHERE username = ?", username).
		Scan(&u.UserId, &u.Username, &u.CreatedAt)
	return u, err
}
