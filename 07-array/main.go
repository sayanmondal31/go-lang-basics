package main

import "fmt"

func main() {
	fmt.Println("welcome to array in go lang")

	// defining array
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list length is: ", len(fruitList))

	var vegList = [3]string{"carrot", "beans", "onion"}

	fmt.Println("vegy list is: ", vegList)
	fmt.Println("vegy list is: ", len(vegList))

}
