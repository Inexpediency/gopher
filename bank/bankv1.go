// package bank provides a parallel secure Bank with a single account

// realisation with 2 channels

package bank

var deposits = make(chan int) // The administration of the contribution
var balances = make(chan int) // Getting a balance

// DepositV1 increases balance on amount
func DepositV1(amount int) {
	deposits <- amount
}

// BalanceV1 returns current balance
func BalanceV1() int {
	return <-balances
}

func teller() {
	var balance int // balance limited in `teller` goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // Starting the control goroutine
}
