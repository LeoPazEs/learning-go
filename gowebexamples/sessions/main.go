package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	session, err := NewSession("338069524db0c8a467044861d9bc2849cf8bd9665cc5f9babe9b77fa97098819", EncodingHex, sha256.New)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	userData := &SessionData{true}
	cookie, err := session.Create(userData)
	fmt.Println(cookie)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	validatedSession, err := session.Check(cookie)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(validatedSession.Authenticated)
	validatedSession, err = session.Check("eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.uv7g8AQZzrEJf9P9y3uRyMOZ8eRbdZpWjLT5dCsG96Q")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(validatedSession.Authenticated)
}
