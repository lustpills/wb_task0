package caching

import (
	"sync"

	"github.com/lustpills/wb_task0/pkg/service"
)

var MyCache *Cache

// simple caching system using map[string]interface{}
type Cache struct {
	services *service.Service
	sync.RWMutex
	items map[string]interface{}
}

// NewCash initiates a caching system
// when cache is initiated it automatically restores orders data from the db
func NewCash(servises *service.Service) *Cache {

	cache := Cache{
		services: servises,
	}

	cache.items, _ = cache.services.Orders.RestoreCache()

	MyCache = &cache

	return &cache
}
