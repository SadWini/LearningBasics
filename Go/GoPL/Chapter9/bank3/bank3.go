package bank3

import "sync"

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	defer mu.Unlock()
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}
