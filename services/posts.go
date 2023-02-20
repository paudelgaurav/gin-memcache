package services

import (
	"encoding/json"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/paudelgaurav/gin-memcache/infrastructure"
	"github.com/paudelgaurav/gin-memcache/models"
	"github.com/paudelgaurav/gin-memcache/repository"
)

func GetPosts() (posts []models.BlogPost, err error) {
	// fetching from cache
	// if available, we will unmarshall and return
	cachedPosts, err := infrastructure.Cache.Get("posts")
	if err == nil {
		posts = toBlogPost(cachedPosts.Value)
		return posts, err
	}
	// fetching from db
	posts, err = repository.GetPosts()
	if err != nil {
		return nil, err
	}
	// storing in cache
	cachedData, err := json.Marshal(posts)
	if err != nil {
		return posts, err
	}
	if err = infrastructure.Cache.Set(&memcache.Item{Key: "posts", Value: cachedData}); err != nil {
		return posts, err
	}
	return posts, err
}

func toBlogPost(data []byte) (posts []models.BlogPost) {
	if err := json.Unmarshal(data, &posts); err != nil {
		return posts
	}
	return posts
}
