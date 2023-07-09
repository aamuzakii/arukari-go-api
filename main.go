package main

import (
	"arukari/initializers"
	"arukari/models"
	"arukari/util"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

		if res.Error != nil {
			errorMsg := res.Error.Error()
			log.Println(errorMsg)
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
			errMsg := tx.Error.Error()
			log.Println(errMsg)
			c.JSON(http.StatusInternalServerError, gin.H{"msg": errMsg})
			return
		}

		c.JSON(http.StatusOK, gin.H{"email": email})
	})

	r.POST("/clock-in", func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		key := []byte("secret")

		if accessToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Authorization token not provided",
			})
			return
		}

		token, err := util.VerifyToken(accessToken, key)

		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusForbidden, gin.H{
				"msg": err.Error(),
			})
			return
		}

		payload := token.Claims.(jwt.MapClaims)

		email := payload["email"].(string)

		employee, err := util.GetEmployee(email)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
			return
		}

		err2 := util.CreateAttendance(employee.ID)

		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err2.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "success add attendance log",
		})
	})

	r.Run()
}
