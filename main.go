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
			"hometown": "Daejeon",
		}
	// dictionary := mydict.Dictionary{}
	// dictionary["name"] = "dasom"
	// fmt.Println(dictionary)

	// value, getError := dictionary.Search("name")
	// if getError != nil {
	// 	fmt.Println("error!!!!", getError)
	// 	return
	// }
	// fmt.Println("value!!!", value)

	addError := dictionary.Add("age", "32")
	if addError != nil {
		fmt.Println("addError", addError)
	} else {
		fmt.Println("add done.", dictionary)
	}

	updateError := dictionary.Update("age", "30")
	if updateError != nil {
		fmt.Println("updateError", updateError)
	} else {
		fmt.Println("update done", dictionary)
	}

	deleteError := dictionary.Delete("age")
	if deleteError != nil {
		fmt.Println("deleteError", deleteError)
	} else {
		fmt.Println("hometown delete done", dictionary)
	}

}
