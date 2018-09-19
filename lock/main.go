// http://www.alexedwards.net/blog/understanding-mutexes
package main

import (
	"fmt"
	"sync"

	bank "github.com/chonlatee/go-get-th/lock/bankAccount"
)

func main() {
	var wg sync.WaitGroup

	for index := 0; index < 5; index++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			bank.Balance.Deposit(100)
			fmt.Printf("Balance after deposit 100: %v \n", bank.Balance.Display())
		}()

		go func() {
			defer wg.Done()
			bank.Balance.Deposit(50)
			fmt.Printf("Balance after deposit 50: %v \n", bank.Balance.Display())
		}()
	}

	wg.Wait()

}
