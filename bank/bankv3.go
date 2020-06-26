// package bank provides a parallel secure Bank with a single account

// realisation with sync.Mutex

package bank

import (
	"fmt"
	"io"
	"sync"
	"time"
)

var (
	billBalance int
	mu          sync.RWMutex // protect the balance
)

// DepositV3 increases balance on amount
func DepositV3(amount int) {
	mu.Lock()
	billBalance += amount
	mu.Unlock()
}

// BalanceV3 returns current balance
func BalanceV3() int {
	mu.RLock()
	defer mu.RUnlock()
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

// Monitor is monitoring bill balance
func Monitor(out io.Writer) int {
	tick := time.Tick(20 * time.Millisecond)

	for {
		select {
		case <- tick:
			mu.RLock()
			fmt.Fprintf(out, "Current balance: %d\n", billBalance)
			mu.RUnlock()
		}
	}
}
