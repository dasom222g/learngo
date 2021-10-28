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

func checkDrink(age int) bool {
	if koreanAge := age + 1; koreanAge < 20 {
		return false
	} else if age < 30 {
		koreanAge = 10
	}
	return true
}

// func checkDrinkSwitch(age int) bool {
// 	switch {
// 	case age < 20:
// 		return false
// 	case age == 20:
// 		return true
// 	case age > 50:
// 		return false
// 	}
// 	return true
// }

func checkDrinkSwitch(age int) bool {
	switch koreanAge := age + 1; koreanAge {
	case 10:
		return false
	case 20:
		return true
	case 50:
		return false
	}
	return true // default
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
	// fmt.Println(totalAdd(1, 2, 3, 4, 5))
	// fmt.Println(checkDrink(50))
	// fmt.Println(checkDrinkSwitch(59))
	// arr := [5]int{1, 2, 3, 4, 5}
	// a := 1
	// b := &a
	// *b = 10
	// fmt.Println(b, *b)
	// fmt.Println(a)

	// arr := []string{"dasom", "kelly"}
	// arr = append(arr, "test")
	// fmt.Println(arr)

	// obj := map[string]bool{
	// 	"isError": false,
	// 	"isValid": false,
	// }

	// for key, value := range obj {
	// 	fmt.Println(key, value)
	// }
	// fmt.Println(obj)

	type person struct {
		name     string
		age      int
		favFoods []string
	}

	favFoods := []string{"pringles", "pasta"}

	personArr := []person{}
	dasom := person{
		name:     "dasom",
		age:      32,
		favFoods: favFoods,
	}
	song := person{
		name:     "song",
		age:      26,
		favFoods: favFoods,
	}
	personArr = append(personArr, dasom, song)
	fmt.Println(personArr)

	for index, person := range personArr {
		fmt.Println(index, person.name)
	}
	fmt.Println(dasom)
}
