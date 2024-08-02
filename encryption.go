package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Example plaintext
	plaintext := "Hello, world!"
	// MD5 Hashing
	md5Hash := md5.Sum([]byte(plaintext))
	fmt.Printf("MD5 Hash: %x\n", md5Hash)

	// SHA-256 Hashing
	sha256Hash := sha256.Sum256([]byte(plaintext))
	fmt.Printf("SHA256 Hash: %x\n", sha256Hash)
}
