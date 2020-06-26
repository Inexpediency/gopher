package main

import (
	"fmt"
	"os"

	"github.com/ythosa/gobih/bank"
)

func main() {
	go bank.Monitor(os.Stdout)

	for i := 0; i < 1000000; i++ {
		bank.DepositV3(1014012)
		bank.WithdrawV3(100)
		bank.BalanceV3()
		bank.DepositV3(100)
		bank.WithdrawV3(100)
	}

	fmt.Println("Balance:", bank.BalanceV3())
}
