// /////////////////////////////////
package main

import "fmt"

func main() {
	a := 12
	b := 6

	isABiggerThanB := a > b

	if isABiggerThanB {
		fmt.Println("bigger")
	} else {
		fmt.Println("smaller")
	}

	// +, -, *, /    ->   int
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	// >, <, ==     ->    bool
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)

}
