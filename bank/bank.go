// bank provides a parallel secure Bank with a single account

package concurrency

var deposits = make(chan int) // The administration of the contribution
var balances = make(chan int) // Getting a balance

// Deposit increases balance on amount
func Deposit(amount int) {
	deposits <- amount
}

// Balance returns current balance
func Balance() int {
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
