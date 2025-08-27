package main

import (
	"fmt"
	"sync"
	"time"
)

/**
go 中并发同样存在线程安全问题，因为 Go 也是使用共享内存让多个
goroutine 之间通信。并且大部分时候为了性能，所以 go 的大多数标准库的数据结构默认是非线程安全的。
*/

//todo  go语言互斥锁示例

// SafeCounter 是一个带有互斥锁的计数器结构体
type SafeCounter struct {
	mu    sync.Mutex // 互斥锁，用于保护 count 的并发访问
	count int        // 计数器值
}

// 增加计数
// GetCount 安全地获取当前计数器值
// defer 确保锁一定会被释放，即使发生 panic
func (c *SafeCounter) Increment() {
	c.mu.Lock()         // 获取互斥锁，阻塞其他 goroutine 访问
	defer c.mu.Unlock() // 延迟释放互斥锁，确保函数返回前解锁
	c.count++           // 返回受保护的计数器值
}

// 获取当前计数
// Increment 安全地增加计数器值
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()         // 获取互斥锁
	defer c.mu.Unlock() // 确保函数返回前解锁
	return c.count      // 增加计数器值
}

func main() {
	//创建SafeCounter实例
	counter := SafeCounter{}

	//使用 WaitGroup 等待所有 goroutine完成
	var wg sync.WaitGroup

	//启动100个goroutine 同时增加计数器
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
			fmt.Println("当前计数 = %d\n", counter.GetCount()) //这行代码如果放在sleep后面，那么打印出来的可能都是10
			time.Sleep(time.Millisecond)
		}()
	}

	//启动10个goroutine同时读取计数器值
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 50) //稍微延迟读取
			fmt.Println("读取器 %d：当前计数 = %d\n", id, counter.GetCount())
		}(i)
	}
	//等待所有 goroutine 完成
	wg.Wait()

	//输出最终计数
	fmt.Println("\nz最终计数：%d\n", counter.GetCount())

}

/**
执行流程详解
1. wg.Add(1) - 增加 WaitGroup 计数器
wg 是一个 sync.WaitGroup 实例，用于等待一组 goroutine 完成

Add(1) 表示要向 WaitGroup 中添加一个待完成的任务（goroutine）

这必须在启动新的 goroutine 之前调用，以确保计数器正确递增

2. go func(id int) { ... }(i) - 启动 Goroutine
go 关键字启动一个新的并发执行流（goroutine）

这是一个匿名函数，接受一个 id 参数

(i) 将循环变量 i 的值传递给 goroutine 作为参数

重要：这里通过参数传递 i 的值，而不是直接在闭包中引用 i，是为了避免常见的并发陷阱（循环变量捕获问题）

3. Goroutine 内部执行流程
a. defer wg.Done()
defer 关键字确保函数返回前执行 wg.Done()

wg.Done() 减少 WaitGroup 的计数器，表示这个 goroutine 已完成

使用 defer 可以保证即使 goroutine 中出现 panic，计数器也会正确递减

b. time.Sleep(time.Millisecond * 50)
让当前 goroutine 暂停执行 50 毫秒

这样做的目的是让增加计数的 goroutine 有更多时间先执行

在实际应用中，这种硬编码的延迟通常不是最佳实践，这里只是为了演示

c. fmt.Printf("读取器 %d：当前计数 = %d\n", id, counter.GetCount())
安全地读取并打印计数器的当前值

注意：这里应该使用 fmt.Printf 而不是 fmt.Println（原代码有误）

counter.GetCount() 会获取互斥锁，读取值，然后释放锁

整体执行流程
主循环迭代 10 次（i 从 0 到 9）

每次迭代：

调用 wg.Add(1) 增加等待计数

启动一个新的 goroutine，并将当前循环变量 i 的值传递给它

每个 goroutine：

立即睡眠 50 毫秒

睡眠结束后，安全地读取计数器值并打印

通过 defer wg.Done() 标记自己完成

主函数中的 wg.Wait() 会阻塞，直到所有 10 个读取器 goroutine 都调用了 wg.Done()

为什么这样设计？
避免竞态条件：通过参数传递循环变量值，而不是直接引用，避免了常见的并发问题

确保资源清理：使用 defer 确保无论 goroutine 如何结束，都会调用 wg.Done()

控制执行顺序：通过睡眠延迟，让写入操作有更多机会先执行（虽然这不是保证）

同步并发任务：使用 WaitGroup 等待所有并发任务完成后再继续

这种模式是 Go 并发编程中的常见做法，特别是在需要启动多个并发任务并等待它们全部完成的场景中。
*/

/**

并发安全原理
互斥锁保护：通过 sync.Mutex 确保同一时间只有一个 goroutine 能访问计数器

原子操作：计数器的增操作和读操作都被保护，不会出现竞态条件

有序访问：锁机制确保了操作的顺序性，避免了数据不一致

执行流程
创建10个写goroutine，每个都会增加计数器

创建10个读goroutine，每个都会读取并打印当前计数

所有goroutine完成后，打印最终计数

可能的问题
代码中的注释有些混乱，方法名和注释不匹配

读取器可能会在计数器还未完全增加时读取，因此可能看到中间值

使用 time.Sleep 只是为了演示，实际应用中应避免这种不确定性

输出结果特点
读取器输出的值可能各不相同，取决于它们执行时计数器的状态

最终计数应该是10（因为有10个写goroutine）

输出顺序是不确定的，因为goroutine的执行顺序由调度器决定

这种模式是Go中处理共享数据的典型方式，通过互斥锁确保数据的一致性和线程安全。
*/
