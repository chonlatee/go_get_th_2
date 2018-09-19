// http://www.alexedwards.net/blog/understanding-mutexes
package main

import (
	"sync"

	bank "github.com/chonlatee/go-get-th/chan/bankAccount"
)

func main() {
	var wg sync.WaitGroup

	sem := make(chan bool, 1)

	for index := 0; index < 5; index++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			sem <- true
			bank.Balance.Deposit(100)
			bank.Balance.Display()
			<-sem
		}()

		go func() {
			defer wg.Done()
			sem <- true
			bank.Balance.Deposit(50)
			bank.Balance.Display()
			<-sem
		}()
	}

	wg.Wait()

}
