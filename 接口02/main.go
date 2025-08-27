package main

import "fmt"

// 接口2
type Account interface {
	getBalance() int
}

type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) getBalance() int {
	return c.balance
}

func main() {
	c := CreditCard{balance: 100, limit: 1000}
	var a Account = &c
	fmt.Println(a.getBalance())
}
