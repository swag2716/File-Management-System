package controllers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {

	// to hash a password with cost factor 14
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panic(err)
	}
	return string(pass)
}
