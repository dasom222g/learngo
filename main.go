package main

import (
	"fmt"

	"github.com/dasom222g/learngo/banking"
)

func main() {
	newAccount := banking.NewAccount("dasom") // account 생성
	newAccount.Deposit(100)
	newAccount.Deposit(300)
	newAccount.GetBalance()

	fmt.Println("GetBalance", newAccount.GetBalance())
}
