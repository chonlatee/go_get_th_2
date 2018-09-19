package bank

import (
	"fmt"
)

type currency struct {
	amount float64
	code   string
}

// Balance export
var Balance = &currency{amount: 0, code: "TH"}

func (c *currency) Deposit(amount float64) {
	fmt.Printf("Deposit: %.2f \n", amount)
	c.amount += amount
}

func (c *currency) Withdraw(amount float64) {
	fmt.Printf("Withdraw: %.2f \n", amount)
	c.amount -= amount
}

func (c *currency) Display() {
	fmt.Printf("Balance now: %.2f%s\n", c.amount, c.code)
}
