package main

import (
	"arukari/initializers"
	"arukari/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.ConnectToDB()
	fmt.Println("mama")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		util.GenerateToken()
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"accessToken": "ey123", "refreshToken": "ey123"})
	})
	r.Run()
}
