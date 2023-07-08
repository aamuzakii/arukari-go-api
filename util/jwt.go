package util

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken() {
	var (
		t *jwt.Token
		// s   string
	)

	key := []byte("secret")

	fmt.Println(key)

	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": "mama.com",
	})

	tokenStr, err := t.SignedString(key)

	fmt.Println(">>>>>>>>>>>>")

	fmt.Println(tokenStr, err)

	res, err2 := VerifyToken(tokenStr, key)

	fmt.Println(res.Claims, err2)
}

func VerifyToken(tokenStr string, key []byte) (*jwt.Token, error) {
	// Define the token verification options
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is as expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the key for validation
		return key, nil
	})

	return token, err
}
