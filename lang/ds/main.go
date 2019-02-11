package main

import (
	"fmt"
	"github.com/GerryLon/learn-go/lang/ds/lru"
)

func traverseCache(cache *lru.CacheDB) {
	cache.Traverse(func(k, v interface{}) bool {
		fmt.Printf("%s:%d\t", k, v)
		return true
	})
	fmt.Println()
}

func LRUTest() {
	cache := lru.New(5) // 缓存最大容量: 5

	// 依次存入 a:1, b:2, ..., e: 5
	for i := 0; i < 5; i++ {
		err := cache.Set(string(rune('a'+i)), i+1)
		if err != nil {
			panic(err)
		}
	}

	// 打印结果: e:5	d:4	c:3	b:2	a:1
	// 也就是最后存储的(最近访问的)在最前边
	traverseCache(cache)

	// 打印结果: f:6	e:5	d:4	c:3	b:2
	// 也就是缓存满后, 再放, 会将最后的元素删除
	cache.Set("f", 6)
	traverseCache(cache)
}

func main() {
	LRUTest()
}