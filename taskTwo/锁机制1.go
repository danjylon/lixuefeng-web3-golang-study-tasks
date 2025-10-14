package main

import (
	"fmt"
	"sync"
)

/*
*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
func main() {
	counter := Counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println("value:", counter.GetValue())
}

type Counter struct {
	sync.Mutex
	value int
}

func (c *Counter) Incr() {
	c.Lock()
	defer c.Unlock()
	c.value++
}
func (c *Counter) GetValue() int {
	c.Lock()
	defer c.Unlock()
	return c.value
}
