package main

import (
	"fmt"
	"strings"
)

func add(a, b int) int {
	defer fmt.Println("It is done.")
	fmt.Println("method==>", a+b)
	return a + b
}

func lenAndUpperName(name string) (length int, upperName string) {
	length = len(name)
	upperName = strings.ToUpper(name)
	return
}

func showArguments(words ...string) {
	fmt.Println(words)
}

// func totalAdd(numbers ...int) int {
// 	sum := 0
// 	for index, number := range numbers {
// 		sum += number
// 		fmt.Println(index, number)
// 	}
// 	return sum
// }

func totalAdd(numbers ...int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func main() {
	// const name string = "dasomi"
	// var name string = "dasomi"
	// name := "dasomi"
	// fmt.Println("name", name)

	// fmt.Println("result", add(10, 20))
	// len, upperName := lenAndUpperName("dasomi")
	// fmt.Println("lenAndUpperName===>", len, upperName)

	// showArguments("a", "b", "c", "d")

	// fmt.Println(totalAdd(7, 8, 3, 4, 5, 6))
	fmt.Println(totalAdd(1, 2, 3, 4, 5))
}
