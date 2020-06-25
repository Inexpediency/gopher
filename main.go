package main

import (
	"fmt"

	"github.com/ythosa/gobih/bank"
)

func main() {
	bank.DepositV3(1014012)
	bank.WithdrawV3(100)
	bank.BalanceV3()
	bank.DepositV3(100)
	bank.WithdrawV3(100)

	fmt.Println(bank.BalanceV3())
}
