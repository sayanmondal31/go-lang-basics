package main

import "fmt"

func main() {
	// fmt.Println("Hello defer") // 1
	// defer fmt.Println("world")  // 3
	// fmt.Println("hello 2")  // 2

	// last in first out   mechanism
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")

	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i) // reversing numbers
	}
}
