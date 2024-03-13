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
func NewCash(servises *service.Service) *Cache {

	cache := Cache{
		services: servises,
		items:    make(map[string]interface{}),
	}

	MyCache = &cache

	return &cache
}
