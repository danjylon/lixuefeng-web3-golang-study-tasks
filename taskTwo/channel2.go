package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func main() {
	ch := make(chan int, 100)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 50)
		}
		close(ch)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case num, ok := <-ch:
				if !ok {
					fmt.Println("channel closed")
					return
				}
				fmt.Println("num: ", num)
			default:
				fmt.Println("暂无数据，等待接收")
				time.Sleep(time.Millisecond * 50)
			}
		}
	}()
	wg.Wait()
}
