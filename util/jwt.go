package util

import (
	"github.com/golang-jwt/jwt/v5"
)

var (
	key []byte
	t   *jwt.Token
	s   string
)

key = []byte("23")
fmt.Println(key)
