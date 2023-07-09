package main

import (
	"arukari/initializers"
	"arukari/models"
	"arukari/util"
	"fmt"
	"log"
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
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/login", func(c *gin.Context) {
		// ambil email & pw dari body
		var loginReq LoginRequest

		c.ShouldBindJSON(&loginReq)

		email := loginReq.Email
		pw := loginReq.Password

		var employee models.Employee

		res := initializers.DB.Where("email = ?", email).First(&employee)

		fmt.Printf("%+v <<< res baru+\n", res) // logging with better info with +

		fmt.Println(res.Error, ">>> is error exist?")

		if res.Error != nil {
			errorMsg := res.Error.Error()
			fmt.Println(errorMsg)
			c.JSON(http.StatusNotFound, gin.H{"msg": errorMsg})
			return
		}
		// cek dengan database
		// kalau benar generate token

		err := util.ComparePassword(employee.Password, pw)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "wrong pwd"})
			return
		}

		fmt.Println(email, "<< email")

		jwtPayload := map[string]interface{}{
			"email": email,
		}

		jwtToken := util.GenerateToken(jwtPayload)

		c.JSON(http.StatusOK, gin.H{"accessToken": jwtToken, "refreshToken": pw})
	})

	r.POST("/register", func(c *gin.Context) {

		var registerReq RegisterRequest

		c.ShouldBindJSON(&registerReq)

		email := registerReq.Email
		pw := registerReq.Password

		hashedPW := util.HashPassword(pw)

		user := models.Employee{
			Name:         "Abdullah Al Muzaki",
			Password:     hashedPW,
			Email:        email,
			DepartmentID: 1,
			Position:     "CTO",
		}

		tx := initializers.DB.Create(&user)

		if tx.Error != nil {
			log.Fatal(tx.Error.Error())
		}

		c.JSON(http.StatusOK, gin.H{"email": email})
	})
	r.Run()
}
