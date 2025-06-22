package main

import "fmt"

func main() {
	var a int = 5
	var b float64 = 4.5
	var c bool = false
	var d string = "text"

	e := 5
	f := 4.5
	g := false
	h := "text"

	six := six()
	fmt.Println(six)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}

func six() int {
	return 6
}
