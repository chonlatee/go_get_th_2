package bank

import (
	"strconv"
)

type currency struct {
	amount float64
	code   string
}

// Balance export
var Balance = &currency{amount: 0, code: "TH"}

func (c *currency) Deposit(amount float64) {
	c.amount += amount
}

func (c *currency) Withdraw(amount float64) {
	c.amount -= amount
}

func (c *currency) Display() string {
	return strconv.FormatFloat(c.amount, 'f', 2, 64) + " " + c.code
}
