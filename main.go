package main

import (
	"arukari/initializers"
	"arukari/models"
	"arukari/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	initializers.ConnectToDB()
	fmt.Println("mama")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		util.GenerateToken()
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/login", func(c *gin.Context) {
		// ambil email & pw dari body
		var loginReq LoginRequest

		c.ShouldBindJSON(&loginReq)

		email := loginReq.Email
		pw := loginReq.Password
		// cek dengan database
		// kalau benar generate token
		c.JSON(http.StatusOK, gin.H{"accessToken": email, "refreshToken": pw})
	})

	r.POST("/register", func(c *gin.Context) {

		var registerReq RegisterRequest

		c.ShouldBindJSON(&registerReq)

		// email := registerReq.Email
		pw := registerReq.Password

		hashedPW := util.HashPassword(pw)

		user := models.Employee{
			Name:         "Abdullah Al Muzaki",
			Password:     hashedPW,
			Email:        "email",
			DepartmentID: 1,
			Position:     "CTO",
		}

		id := initializers.DB.Create(&user)

		c.JSON(http.StatusOK, gin.H{"email": id, "refreshToken": pw})
	})
	r.Run()
}
