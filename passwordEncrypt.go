package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Password to be hashed
	password := "MySecuredPassword"

	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Output the hashed password
	fmt.Println("Original Password:", password)
	fmt.Println("Hashed Password:", string(hashedPassword))

	// Comparing a password attempt with the hashed password
	passwordAttempt := "MySecuredPassword"
	//passwordAttempt := "WrongPassword"
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(passwordAttempt))
	if err == nil {
		fmt.Println("Password Matched!")
	} else {
		fmt.Println("Password did not match!")
	}
}
