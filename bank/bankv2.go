// pacakge bank provides a parallel secure Bank with a single account

// realisation with semaphore

package bank

var (
	sema    = make(chan struct{}, 1)
	balance int
)

// DepositV2 increases balance on amount
func DepositV2(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}

// BalanceV2 returns current balance
func BalanceV2() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
