package main

import (
	"fmt"
)

// 接口1
type PayMethod interface {
	Account
	Pay(amount int) bool
}

type Account interface {
	GetBalance() int
}

type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Println("信用卡支付成功：%d\n", amount)
		return true
	}
	fmt.Println("信用卡支付失败：超出额度")
	return false
}

func (c *CreditCard) GetBalance() int {
	return c.balance
}

type DebitCard struct {
	balance int
}

func (d *DebitCard) Pay(amount int) bool {
	if d.balance >= amount {
		d.balance -= amount
		fmt.Println("借记卡支付成功%d\n", amount)
		return true
	}
	fmt.Println("借记卡支付失败：余额不足")
	return false
}

func (d *DebitCard) GetBalance() int {
	return d.balance
}

func purchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Println("购买成功，剩余余额：%d\n", p.GetBalance)
	} else {
		fmt.Println("购买失效")
	}
}

func main() {
	creditCard := &CreditCard{balance: 0, limit: 100}
	debitCard := &DebitCard{balance: 500}

	fmt.Println("使用信用卡购买；")
	purchaseItem(creditCard, 800)

	fmt.Println("\n使用借记卡购买；")
	purchaseItem(debitCard, 300)

	fmt.Println("\n再次使用借记卡购买；")
	purchaseItem(debitCard, 400)
}
