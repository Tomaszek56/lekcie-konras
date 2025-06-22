package main

import (
	"fmt"
)

func checkNumberRange(number int) {
	if number < 10 {
		fmt.Println("Outside the range")
	} else if number < 20 {
		fmt.Println("It is in the range.")
		if number >= 10 && number <= 13 {
			fmt.Println("Its range beginning.")
		} else if number >= 14 && number <= 17 {
			fmt.Println("Its range middle.")
		} else {
			fmt.Println("Its range ending.")
		}
	} else {
		fmt.Println("The number is outside the range:", number)
	}
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	fmt.Println()
	checkNumberRange(number)
}
