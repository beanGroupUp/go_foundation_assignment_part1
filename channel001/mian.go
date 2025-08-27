package main

/**
channel 是 Go 中定义的一种类型，专门用来在多个 goroutine 之间通信的线程安全的数据结构。

可以在一个 goroutine 中向一个 channel 中发送数据，从另外一个 goroutine 中接收数据。

channel 类似队列，满足先进先出原则。

定义方式：
// 仅声明
var <channel_name> chan <type_name>

// 初始化
<channel_name> := make(chan <type_name>)

// 初始化有缓冲的channel
<channel_name> := make(chan <type_name>, 3)

channel 的三种操作：发送数据，接收数据，以及关闭通道。

声明方式：
// 发送数据
channel_name <- variable_name_or_value

// 接收数据
value_name, ok_flag := <- channel_name
value_name := <- channel_name

// 关闭channel
close(channel_name)

channel 还有两个变种，可以把 channel 作为参数传递时，限制 channel 在函数或方法中能够执行的操作。

声明方式：

//仅发送数据
func <method_name>(<channel_name> chan <- <type>)

//仅接收数据
func <method_name>(<channel_name> <-chan <type>)

*/

//值接收
