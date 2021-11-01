package main

import (
	"fmt"

	"github.com/dasom222g/learngo/banking"
)

func main() {
	newAccount := banking.NewAccount("dasom") // account 생성
	newAccount.Deposit(100)
	error := newAccount.Withdraw(500)
	if error != nil {
		// log.Fatal(error)
		fmt.Println(error)
	}

	fmt.Println(newAccount)
}
