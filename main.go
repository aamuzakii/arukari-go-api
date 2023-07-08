package main

import (
	"arukari/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("mama")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		util.PrintToken()
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.Run()
}