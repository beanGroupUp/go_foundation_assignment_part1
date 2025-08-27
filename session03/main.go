package main

import "fmt"

func main() {
	//终端for循环
	/*	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}*/

	/*switch i := 1; i {
	case 1:
		fmt.Println("1")
		if i == 1 {
			break
		}
		fmt.Println("1等于1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}*/

	//中断select
	//如果case中处于阻塞状态，那么整个select都处于阻塞状态
	//select从左到右，从上到下，有default执行default
	/*	select {
		case <-time.After(2 * time.Second):
			fmt.Println("过了两秒")
		case <-time.After(1 * time.Second):
			fmt.Println("过了一秒")
			if true {
				break
			}
			fmt.Println("break之后")
		}
	*/

	//跳出for循环和select循环
	/*for {
			select {
			case <-time.After(time.Second * 3):
				fmt.Println("过了3秒")
			case <-time.After(time.Second * 1):
				fmt.Println("过了1秒")
				//跳出for循环
				goto exit
				if true {
					//跳出select
					break
				}
				fmt.Println("break 之后")
			}
		}

	exit:
		fmt.Println("终端")*/

	/*	for i := 0; i <= 3; i++ {
				fmt.Println("i=", i)
				for j := 0; j < 10; j++ {
					fmt.Println("j=", j)
					if j == 5 {
						break
					}
					goto exit
				}
			}
		exit:
			fmt.Println("终端")*/

	//适用标记
outter:
	for i := 1; i <= 3; i++ {
		fmt.Println("i=", i)
		for j := 0; j < 3; j++ {
			fmt.Println("j=", j)
			break outter
		}
	}
}
