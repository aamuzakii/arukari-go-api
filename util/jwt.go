package util

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload map[string]interface{}) string {
	var (
		t *jwt.Token
		// s   string
	)

	key := []byte("secret")

	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": payload["email"],
	})

	tokenStr, err := t.SignedString(key)

	if err != nil {
		log.Fatal(err.Error())
	}

	// klaim := res.Claims.(jwt.MapClaims)

	return tokenStr

}

func VerifyToken(tokenStr string, params ...[]byte) (*jwt.Token, error) {
	key := []byte("secret")
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
