package caching

import (
	"sync"

	"github.com/lustpills/wb_task0/pkg/service"
)

var MyCache *Cache

type Cache struct {
	services *service.Service
	sync.RWMutex
	items map[string]interface{}
}

// type Item struct {
// 	Value interface{}
// }

func NewCash(servises *service.Service) *Cache {

	// инициализируем карту(map) в паре ключ(string)/значение(Item)
	//items := make(map[string]Item)

	cache := Cache{
		services: servises,
	}

	cache.items, _ = cache.services.Orders.RestoreCache()

	MyCache = &cache

	return &cache
}
