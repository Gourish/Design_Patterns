package main

import "fmt"

type evictionAlgo interface {
	evict(c *cache)
}

type lfu struct{}

func (lfu *lfu) evict(c *cache) {
	fmt.Println("evicting through lfu")
}

type lru struct{}

func (lfu *lru) evict(c *cache) {
	fmt.Println("evicting through lru")
}

type fifo struct{}

func (fifo *fifo) evict(c *cache) {
	fmt.Println("evicting through fifo")
}
