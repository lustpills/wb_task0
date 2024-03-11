package caching

func (c *Cache) Set(key string, value interface{}) {

	c.Lock()

	defer c.Unlock()

	c.items[key] = value

}

func (c *Cache) Get(key string) (interface{}, bool) {

	c.RLock()

	defer c.RUnlock()

	item, found := c.items[key]

	// ключ не найден
	if !found {
		return nil, false
	}
	// Проверка на установку времени истечения, в противном случае он бессрочный

	return item, true
}

func (c *Cache) Restore() error {
	//restored_cache, err := c.services.Orders.RestoreCache()

	//c.items = restored_cache
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return nil
}
