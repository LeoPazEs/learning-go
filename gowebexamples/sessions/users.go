package main

import (
	"crypto/subtle"
	"fmt"
	"sync"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Check in constant time if password is correct.
func (u *User) CheckPassword(password string) bool {
	return subtle.ConstantTimeCompare([]byte(u.Password), []byte(password)) == 1
}

type Users struct {
	mu    sync.Mutex
	Users map[string]*User
}

type (
	ErrUserAlreadyExists  string
	ErrUserAuthentication string
)

func (e ErrUserAlreadyExists) Error() string {
	return fmt.Sprintf("User with username %s already exists.", string(e))
}

func (e ErrUserAuthentication) Error() string {
	return string(e)
}

// Creates User
func (us *Users) AddUser(u *User) (*User, error) {
	us.mu.Lock()
	defer us.mu.Unlock()
	_, exists := us.Users[u.Username]
	if exists {
		return nil, ErrUserAlreadyExists(u.Username)
	}
	us.Users[u.Username] = u
	return u, nil
}

// Get User
func (us *Users) GetUser(username string) (*User, bool) {
	us.mu.Lock()
	defer us.mu.Unlock()
	userFound, exists := us.Users[username]
	return userFound, exists
}

// Authenticate User
func (us *Users) Authenticate(username string, password string) (bool, error) {
	user, exists := us.GetUser(username)
	if !exists {
		return false, ErrUserAuthentication("Username or Password incorrect.")
	}
	if !user.CheckPassword(password) {
		return false, ErrUserAuthentication("Username or Password incorrect.")
	}
	return true, nil
}
