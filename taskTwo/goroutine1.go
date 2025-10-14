package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

/*
*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func main() {
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				fmt.Println("偶数", i)
			}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 != 0 {
				fmt.Println("奇数", i)
			}
		}
	}()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
