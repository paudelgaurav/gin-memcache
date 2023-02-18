package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	serveApplication()
}

func serveApplication() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "ping"})
	})
	r.Run()
}
