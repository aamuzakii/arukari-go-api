package util

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte("23")
	fmt.Println(key)

}
