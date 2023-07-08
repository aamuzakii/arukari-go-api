package util

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func PrintToken() {
	var (
		t *jwt.Token
		// s   string
	)

	key := []byte{'H', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd', '!'}

	fmt.Println(key)

	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": "mama.com",
	})

	tokenStr, err := t.SignedString(key)

	fmt.Println(">>>>>>>>>>>>")

	fmt.Println(tokenStr, err)
}
