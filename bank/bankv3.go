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
	billBalance += amount
	mu.Unlock()
}

// BalanceV3 returns current balance
func BalanceV3() int {
	mu.Lock()
	defer mu.Unlock()
	return billBalance
}

// WithdrawV3 reduces the billBalance by the passed value
func WithdrawV3(amount int) bool {
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}

	return true
}

// deposit requires Mutex.Lock()
func deposit(amount int) {
	balance += amount
}
