package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/sample", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "content",
		})
	})
	r.Run()
}