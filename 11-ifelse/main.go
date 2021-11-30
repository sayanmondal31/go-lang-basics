package main

import "fmt"

func main() {
	fmt.Println("if else is necessery")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "not regular user"
	} else if loginCount < 15 {
		result = "medium regular user"
	} else {
		result = "regular user"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}

	// if err != nil {

	// }
}
