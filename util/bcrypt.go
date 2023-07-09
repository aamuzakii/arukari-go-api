package util

import (
	"log"
	// "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(result)
}

func ComparePassword(hashedPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}
