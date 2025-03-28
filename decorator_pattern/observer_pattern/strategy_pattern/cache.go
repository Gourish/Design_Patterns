package main

type cache struct {
	data     map[string]string
	eviction evictionAlgo
}

func initCache(e evictionAlgo) *cache {
	data := make(map[string]string)
	return &cache{
		data:     data,
		eviction: e,
	}
}

func (c *cache) evictData() {
	c.eviction.evict(c)
}

func (c *cache) addData(key, val string) {
	c.data[key] = val
}

func (c *cache) changeEviction(e evictionAlgo) {
	c.eviction = e
}
