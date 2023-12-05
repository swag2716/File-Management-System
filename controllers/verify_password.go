package controllers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(providedPassword string, userPassword string) (bool, string) {

	// to compare the provided password and the stored user password
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))

	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintln("password is incorrect")
		check = false
	}
	return check, msg
}
