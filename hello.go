package main

import "fmt"

func main() {
	a := 6
	b := 5
	c := 4
	d := 3

	fmt.Println(a, b, c, d)
	sumab := add(a, b)
	sumcd := add(c, d)
	fmt.Println(sumab)
	fmt.Println(sumcd)
	fmt.Println(sub(sumab, sumcd))
	fmt.Println(mult(sumab, sumcd))
}

func add(x int, y int) int {
	return x + y
}
func sub(x int, y int) int {
	return x - y
}
func mult(x int, y int) int {
	return x * y
}
