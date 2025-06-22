package main

import "fmt"

func main() {
	a := 5

	fmt.Println(a)

	a = 3

	fmt.Println(a)

	// a++ -> a = a + 1
	a = a + 1
	fmt.Println(a)

	for i := 0; i < 5; i++ {
		fmt.Println(i, "Tomek")
	}
}
