package main

import (
	"arukari/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string
	Department string
	Position   string
}

func main() {
	fmt.Println("mama")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		util.GenerateToken()
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.Run()
}
