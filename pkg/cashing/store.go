package caching

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

func (c *Cache) Restore() error {
	return nil
}
