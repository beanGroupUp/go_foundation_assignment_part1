package main

import "fmt"

func main() {
	//for循环方式1
	/*	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	//for循环方式2
	/*	b := 1
		for b < 10 {
			fmt.Println(b)
			b++
			if b > 5 {
				break
			}
		}*/

	//无限循环方式3
	/*ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	var started bool
	var stopped atomic.Bool
	for {
		if !started {
			started = true
			go func() {
				select {
				case <-ctx.Done():
					fmt.Println("ctx done")
					stopped.Store(true)
					return
				}
			}()
		}
		fmt.Println("main")
		if stopped.Load() {
			break
		}
	}*/

	//遍历数组
	/*	var a [10]string
		a[0] = "hello"
		for i := range a {
			fmt.Println("当前下标：", i)
		}

		for i, e := range a {
			fmt.Println("a[", i, "]=", e)
		}*/

	m := make(map[string]string)
	m["b"] = "Helloo,b"
	m["a"] = "Helloo,a"
	m["c"] = "Helloo,c"
	for i, _ := range m {
		fmt.Println("当前key:", i)
	}
	for k, v := range m {
		fmt.Println("k = ", k, "v = ", v)
	}
}
