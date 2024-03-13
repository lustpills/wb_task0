package caching

import "log"

// Set adds an item to the cache
func (c *Cache) Set(key string, value interface{}) {

	c.Lock()

	defer c.Unlock()

	c.items[key] = value

}

// Get retuns an item from cache by a string key
func (c *Cache) Get(key string) (interface{}, bool) {

	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]

	if !found {
		return nil, false
	}
	return item, true
}

// put orders from a db to cache
func (c *Cache) Restore() error {
	items, err := c.services.Orders.RestoreCache()
	if err != nil {
		log.Println("Error occured when restoring cache: ", err)
		return err
	}
	c.items = items
	return nil
}
