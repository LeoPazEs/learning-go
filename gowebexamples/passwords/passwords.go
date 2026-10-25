package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	var password string
	_, err := fmt.Scanln(&password)
	if err != nil {
		log.Println(err)
		return
	}
	hash := HashPassword(password)
	fmt.Println(hash)

	_, _ = fmt.Scanln(&password)
	fmt.Println(CheckPasswordHash(password, hash))
}
