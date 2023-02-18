package services

import (
	"encoding/json"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/paudelgaurav/gin-memcache/infrastructure"
	"github.com/paudelgaurav/gin-memcache/models"
	"github.com/paudelgaurav/gin-memcache/repository"
)

func GetPosts() (posts []models.BlogPost, err error) {
	cachedPosts, err := infrastructure.Cache.Get("posts")
	if err == nil {
		posts = toBlogPost(cachedPosts.Value)
		return
	}
	// fetching fro, db and setting up cache
	posts, err = repository.GetPosts()
	if err != nil {
		return
	}
	// storing in cache
	cachedData, err := json.Marshal(posts)
	if err != nil {
		return
	}
	if err = infrastructure.Cache.Set(&memcache.Item{Key: "posts", Value: cachedData}); err != nil {
		return
	}
	return
}

func toBlogPost(data []byte) (posts []models.BlogPost) {
	if err := json.Unmarshal(data, &posts); err != nil {
		return
	}
	return posts
}
