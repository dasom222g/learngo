package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello")
	// const name string = "dasomi"
	// var name string = "dasomi"
	name := "dasomi"

	fmt.Println("name " + name)
	fmt.Println(add(10, 20))
}
