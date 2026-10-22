package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	session, err := NewSession("BC8bDwK+JlLpR1gPHBZPU962UoL8A2ARC36vC4csjac=", sha256.New)
	if err != nil {
		fmt.Println(err)
	}
	userData := &SessionData{true}
	cookie, err := session.Create(userData)
	if err != nil {
		fmt.Println(err)
	}
	validatedSession, err := session.Check(cookie)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(validatedSession)
}
