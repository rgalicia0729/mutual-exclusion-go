package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Name    string
	Balance float64
}

func transfer(amount float64, source, dest *Account) {
	if source.Balance < amount {
		fmt.Printf("Error %+v, %+v\n", source, dest)
		return
	}

	time.Sleep(time.Second)

	dest.Balance += amount
	source.Balance -= amount

	fmt.Printf("Exito %+v, %+v\n", source, dest)
}

func main() {
	sofia := Account{
		Name:    "Sofía Galicia",
		Balance: 500,
	}

	david := Account{
		Name:    "David Galicia",
		Balance: 900,
	}

	amountTransfers := []float64{300.00, 200.00, 100.00}

	var wg sync.WaitGroup
	wg.Add(len(amountTransfers))

	var mu sync.Mutex
	for _, value := range amountTransfers {
		go func(amount float64) {
			mu.Lock()
			transfer(amount, &sofia, &david)
			mu.Unlock()
			wg.Done()
		}(value)
	}

	wg.Wait()
}
