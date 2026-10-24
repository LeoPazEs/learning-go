package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	UsersDB      = &Users{Users: make(map[string]*User)}
	SessionMaker *Session
)

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user *User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := UsersDB.AddUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Actually slower than just do a marshal of the entire thing
	safeBytes, err := json.Marshal(u.Username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	_, err = fmt.Fprintf(w, `{"username":%s}`, safeBytes)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user *User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = UsersDB.Authenticate(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	sessionValue, err := SessionMaker.Create(&SessionData{true})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionValue,
		Path:     "/",                            // Available site-wide
		Domain:   "localhost",                    // Scoped to your domain
		Expires:  time.Now().Add(24 * time.Hour), // Hard deadline
		MaxAge:   86400,                          // Time-to-live in seconds (takes precedence over Expires)
		HttpOnly: true,                           // Blocks client-side JS access (Prevents XSS)
		SameSite: http.SameSiteLaxMode,           // Protection against CSRF attacks
	}
	http.SetCookie(w, cookie)
}

func secretHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprint(w, `{"secret": "Coding still matters!"}`)
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	SessionMaker, err = NewSession("338069524db0c8a467044861d9bc2849cf8bd9665cc5f9babe9b77fa97098819", EncodingHex, sha256.New)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	http.HandleFunc("POST /users", Chain(createUserHandler, ContentTypeCheck("application/json")))
	http.HandleFunc("POST /login", Chain(loginUserHandler, ContentTypeCheck("application/json")))
	http.HandleFunc("GET /secret", Chain(secretHandler, LoginCheck("session_token"), ContentTypeCheck("application/json")))

	err = http.ListenAndServe(":8080", nil)
	panic(err)
}
