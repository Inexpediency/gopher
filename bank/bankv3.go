// pacakge bank provides a parallel secure Bank with a single account

// realisation with sync.Mutex

package bank

import "sync"

var (
	billBalance int
	mu          sync.Mutex // protect the balance
)

// DepositV3 increases balance on amount
func DepositV3(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

// BalanceV3 returns current balance
func BalanceV3() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
