package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

/*
*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信
*/
func main() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 200)
		}
		close(ch)
	}()
	go func() {
		for {
			num, ok := <-ch
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println("num=", num)
		}
	}()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
