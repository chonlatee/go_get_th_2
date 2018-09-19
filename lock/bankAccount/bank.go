package bank

import (
	"strconv"
	"sync"
)

type currency struct {
	mu     sync.Mutex
	amount float64
	code   string
}

// Balance export
var Balance = &currency{amount: 0, code: "TH"}

func (c *currency) Deposit(amount float64) {
	c.mu.Lock()
	c.amount += amount
	c.mu.Unlock()
}

func (c *currency) Display() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return strconv.FormatFloat(c.amount, 'f', 2, 64) + " " + c.code
}
