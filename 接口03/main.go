package main

import "fmt"

// 接口3
type PayMethod interface {
	Pay(int)
}

type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) Pay(amount int) {
	if c.balance < amount {
		fmt.Println("余额不足")
		return
	}
	c.balance -= amount
}

func anyParam(param interface{}) {
	fmt.Println("param", param)
}

func main() {
	c := CreditCard{balance: 100, limit: 1000}
	c.Pay(10)
	var a PayMethod = &c
	anyParam(a)

	var b PayMethod = &c
	fmt.Println("b:", b)

	anyParam(c)
	anyParam(1)
	anyParam("123")
	anyParam(a)
}
