package main

import (
	// "github.com/dasom222g/learngo/banking"
	"errors"
	"fmt"
	"net/http"

	"github.com/dasom222g/learngo/mydict"
)

type requestResult struct {
	url    string
	status string
}

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
		// fmt.Println("addError", addError)
	} else {
		// fmt.Println("add done.", dictionary)
	}

	updateError := dictionary.Update("age", "30")
	if updateError != nil {
		// fmt.Println("updateError", updateError)
	} else {
		// fmt.Println("update done", dictionary)
	}

	deleteError := dictionary.Delete("age")
	if deleteError != nil {
		// fmt.Println("deleteError", deleteError)
	} else {
		// fmt.Println("hometown delete done", dictionary)
	}

	/*  URL CHECKER & GO ROUTINES */

	resultMap := make(map[string]string)
	ch := make(chan requestResult)
	/*
		[
			{"https://www.airbnb.com/" : "OK"},
			{"https://www.google.com/" : "OK"},
			...
		]
	*/

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go checkUrl(url, ch)
	}

	for i := 0; i < len(urls); i++ {
		message := <-ch
		resultMap[message.url] = message.status
	}

	for url, status := range resultMap {
		fmt.Println(url, status)
	}

}

var errorRequestFailed = errors.New("Failed to get data.")

func checkUrl(url string, ch chan<- requestResult) {
	response, err := http.Get(url)
	status := "OK"
	if err != nil || response.StatusCode >= 400 {
		status = "FAILED"
	}
	ch <- requestResult{url: url, status: status}
}
