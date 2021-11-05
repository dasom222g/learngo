package main

import (
	"fmt"

	// "github.com/dasom222g/learngo/banking"
	"github.com/dasom222g/learngo/mydict"
)

func main() {
	// account
	/*
		newAccount := banking.NewAccount("dasom") // account 생성
		newAccount.Deposit(100)
		error := newAccount.Withdraw(500)
		if error != nil {
			// log.Fatal(error)
			fmt.Println(error)
		}

		fmt.Println(newAccount)
	*/
	dictionary :=
		mydict.Dictionary{
			"name":     "kelly",
			"homeTown": "Daejeon",
		}
	// dictionary := mydict.Dictionary{}
	// dictionary["name"] = "dasom"
	// fmt.Println(dictionary)

	value, error := dictionary.Search("age")
	if error != nil {
		fmt.Println("error!!!!", error)
		return
	}
	fmt.Println("value!!!", value)
}
