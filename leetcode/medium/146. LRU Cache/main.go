// https://leetcode.com/problems/lru-cache/description/
package main

import "fmt"

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)           // кэш: {1=1}
	lru.Put(2, 2)           // кэш: {1=1, 2=2}
	fmt.Println(lru.Get(1)) // 1
	lru.Put(3, 3)           // вытесняем ключ 2, кэш: {1=1, 3=3}
	fmt.Println(lru.Get(2)) // -1
	lru.Put(4, 4)           // вытесняем ключ 1, кэш: {4=4, 3=3}
	fmt.Println(lru.Get(1)) // -1
	fmt.Println(lru.Get(3)) // 3
	fmt.Println(lru.Get(4)) // 4
}

type LRUCache struct {
	// TODO: решить
}

func Constructor(capacity int) LRUCache {
	return LRUCache{}
}

func (this *LRUCache) Get(key int) int {
	return 0
}

func (this *LRUCache) Put(key int, value int) {
}
