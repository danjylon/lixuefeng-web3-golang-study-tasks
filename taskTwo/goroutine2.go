package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"strings"
	"time"
)

/*
*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func main() {
	funcs := []func(){foo, bar, baz, qux}
	scheduler(funcs)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}

func scheduler(funcs []func()) {
	for _, fun := range funcs {
		go func() {
			start := time.Now().UnixMilli()
			// 获取函数指针
			funcPtr := reflect.ValueOf(fun).Pointer()
			// 获取函数对象
			funcObj := runtime.FuncForPC(funcPtr)
			// 提取纯函数名（去掉包路径）
			funcName := strings.TrimPrefix(funcObj.Name(), "main.")
			fun()
			end := time.Now().UnixMilli()
			fmt.Printf("函数%s运行时间：%d毫秒\n", funcName, end-start)
		}()
	}
}

func foo() {
	num := rand.Intn(1000)
	time.Sleep(time.Duration(num) * time.Millisecond)
}

func bar() {
	num := rand.Intn(1000)
	time.Sleep(time.Duration(num) * time.Millisecond)
}
func baz() {
	num := rand.Intn(1000)
	time.Sleep(time.Duration(num) * time.Millisecond)
}
func qux() {
	num := rand.Intn(1000)
	time.Sleep(time.Duration(num) * time.Millisecond)
}
