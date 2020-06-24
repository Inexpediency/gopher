// pacakge bank provides a parallel secure Bank with a single account

// realisation with semaphore

package bank

var (
	sema    = make(chan struct{}, 1) // A binary semaphore to
	balance int                      // protect the balance
)

// DepositV2 increases balance on amount
func DepositV2(amount int) {
	sema <- struct{}{} // The capture of the marker
	balance = balance + amount
	<-sema // The release of the marker
}

// BalanceV2 returns current balance
func BalanceV2() int {
	sema <- struct{}{} // The capture of the marker
	b := balance
	<-sema // The release of the marker
	return b
}
