package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-memcache/infrastructure"
	"github.com/paudelgaurav/gin-memcache/models"
)

func CreateBlogPost(c *gin.Context) {
	posts := []models.BlogPost{
		{
			Title:   "First Title",
			Content: "First content",
		},
		{
			Title:   "First Title",
			Content: "First content",
		}, {
			Title:   "First Title",
			Content: "First content",
		}, {
			Title:   "First Title",
			Content: "First content",
		}, {
			Title:   "First Title",
			Content: "First content",
		}, {
			Title:   "First Title",
			Content: "First content",
		}, {
			Title:   "First Title",
			Content: "First content",
		}, {
			Title:   "First Title",
			Content: "First content",
		}, {
			Title:   "First Title",
			Content: "First content",
		},
	}

	if err := infrastructure.DB.Create(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "blog posts created"})
}

func GetBlogPosts(c *gin.Context) {
	var posts []models.BlogPost
	if err := infrastructure.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while fetching blog posts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts})
}
