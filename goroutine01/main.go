package main

import (
	"fmt"
	"time"
)

/**
go 支持并发的方式，就是通过 goroutine 和 channel 提供的简洁且高效的方式实现的。

goroutine 是轻量线程，创建一个 goroutine 所需的资源开销很小，所以可以创建非常多的 goroutine 来并发工作。

它们是由 Go 运行时调度的。调度过程就是 Go 运行时把 goroutine 任务分配给 CPU 执行的过程。

但是 goroutine 不是通常理解的线程，线程是操作系统调度的。

在 Go 中，想让某个任务并发或者异步执行，只需把任务封装为一个函数或闭包，交给 goroutine 执行即可。

声明方式 1，把方法或函数交给 goroutine 执行
*/

/**
声明方式 1，把方法或函数交给 goroutine 执行：
go <method_name>(<method_params>...)
*/

/**
声明方式 2，把闭包交给 goroutine 执行：
go func(<method_params>...){
    <statement_or_expression>
    ...
}(<params>...)
*/

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// 启动第一个goroutine：使用匿名函数闭包
	go func() {
		fmt.Println("run goroutine in closure01")
	}()
	go func(msg string) {
		fmt.Println(msg)
	}("run goroutine in closure02")
	say("run goroutine in closure03")
	go say("run goroutine in closure04") //这种写法，日志可能打印不出来，因为程序不会等他执行完
	time.Sleep(10000 * time.Second)      //加上等待时间就可以打印出来了

	// 注意：main函数不会等待goroutines完成
	// 如果main退出，所有goroutines会被终止
	// 因此可能需要使用同步机制（如WaitGroup）来等待goroutines完成
}
