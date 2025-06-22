package main

import "fmt"

func CheckAge(age int) {
	if age > 0 {
		fmt.Println("Baby")
	} else if age >= 5 {
		fmt.Println("Kid")
	} else if age >= 10 {
		fmt.Println("Teenager")
	} else if age >= 19 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Age incorrect")
	}
}

func main() {
	CheckAge(11)
}
