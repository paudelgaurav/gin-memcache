package repository

import (
	"time"

	"github.com/paudelgaurav/gin-memcache/infrastructure"
	"github.com/paudelgaurav/gin-memcache/models"
)

func GetPosts() (posts []models.BlogPost, err error) {
	time.Sleep(time.Second * 2)
	return posts, infrastructure.DB.Find(&posts).Error
}
