package main

import "fmt"

func main() {
	fmt.Println("welcome to functions")
	greet()

	result := adder(3, 5)
	fmt.Println("result is : ", result)

	proResult, message := proAdder(5, 9, 8, 7, 9, 89)
	fmt.Println("result is : ", proResult)
	fmt.Println("message is : ", message)
}

// func sum( n1 int,n2)  {
// 	return n1 + n2;
// }

func greet() {
	fmt.Println("hello, নমস্কার , नमस्ते ")
}

// func greet() {
// 	fmt.Println("hello, নমস্কার , नमस्ते ")
// }

func adder(valOne int, valTwo int) int { // we need to add type what we especting
	// from functions --> this is called as signature
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) { // "..." veriatic functions
	total := 0
	for _, val := range values {
		total += val

	}

	return total, "hi pro result"
}
