/*
*  如果我们不需要chanel通信，只想安全的访问一个共享变量
*  使用互斥锁来解决
*  sync.Mutex互斥类型及两个方法：Lock() Unlock()
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发安全使用
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++ // Lock后同一时刻只有一个goroutine可以访问c.v
}

// Value 返回给定key的计数器的当前值
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Inc("key")
	}

	time.Sleep(time.Second)
	fmt.Println(c.value("key"))
}
