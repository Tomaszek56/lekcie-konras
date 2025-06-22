// ////////////////////////////////////
package main

import "fmt"

// if <bool> {
// } else if <bool2> {
// } else if <bool3> {
// } else {
// }

func main() {
	isTrue := false
	isTrue2 := false
	isTrue3 := false
	isTrue4 := true
	isTrue5 := true

	if isTrue {
		fmt.Println("before 1")
	} else if isTrue2 {
		fmt.Println("before 2")
	} else if isTrue3 {
		fmt.Println("before 3")
	} else if isTrue4 {
		fmt.Println("before 4")
	} else if isTrue5 {
		fmt.Println("before 5")
	} else {
		fmt.Println("before else")
	}
	//ik heb mijn hanmden niet gewassen na het plassen
	fmt.Println("after")
}
