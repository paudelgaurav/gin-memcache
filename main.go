package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-memcache/controllers"
	"github.com/paudelgaurav/gin-memcache/infrastructure"
	"github.com/paudelgaurav/gin-memcache/models"
)

func main() {
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	infrastructure.SetUpMemCache()
	infrastructure.SetUpDB()
	if err := infrastructure.DB.AutoMigrate(&models.BlogPost{}); err != nil {
		log.Fatal("BlogPost migrate err", err)
	}
}

func serveApplication() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "ping"})
	})
	r.POST("/posts", controllers.CreateBlogPost)
	r.GET("/posts", controllers.GetBlogPosts)
	r.POST("/clear", func(c *gin.Context) {
		infrastructure.Cache.DeleteAll()
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	})
	r.Run()
}
