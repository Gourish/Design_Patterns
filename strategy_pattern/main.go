package main

func main() {
	lfu := &lfu{}
	cache := initCache(lfu)
	cache.addData("1", "hello")
	cache.evictData()
	lru := &lru{}
	cache.changeEviction(lru)
	cache.addData("2", "yellow")
	cache.evictData()
}
