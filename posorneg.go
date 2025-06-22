package main

import (
	"fmt"
)

func main() {
	for i := -5; i <= 5; i++ {
		fmt.Println(i)
		positveornegative(i)
	}
}
func positveornegative(number int) {
	if number < 0 {
		fmt.Println("negative")
	} else if number > 0 {
		fmt.Println("positive")
	} else if number == 0 {
		fmt.Println("zero")
	}

}
