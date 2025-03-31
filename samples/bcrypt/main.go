package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	bcryptStuffs()
}

func bcryptStuffs() {
	initalPasswd := "1"

	// Generating the hashed password
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(initalPasswd), 10)
	if err != nil {
		log.Fatal("Unable to generate password: ", err)
	}
	fmt.Println(string(hashedPasswd))

	// Checking compliance of the provided password with the original one.
	providedPassword := "Password31551"
	unableToLogin := bcrypt.CompareHashAndPassword(hashedPasswd, []byte(providedPassword))
	if unableToLogin != nil {
		log.Fatal("Unable to login: ", unableToLogin)
	}
	fmt.Println("Passwords match. Welcome.")
}
