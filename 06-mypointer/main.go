package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int // a pointer which stores integer value

	fmt.Println("value of pointer is ", ptr)

	fmt.Println("memory address  is ", &ptr)

	myNumber := 23

	var ptr2 = &myNumber // reference means & || create a pointer and referencing a value

	fmt.Println("Memory address ptr2 is ", ptr2)
	fmt.Println("Value of ptr2 is ", *ptr2)

	*ptr2 = *ptr2 * *ptr2

	fmt.Println("New value is: ", *ptr2)    //with pointer
	fmt.Println("New value is: ", myNumber) // with actual values
}
