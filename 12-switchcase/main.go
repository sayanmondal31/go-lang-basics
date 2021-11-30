package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and cas in golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("can move 2 spots")
	case 3:
		fmt.Println("can move 3 spots")
	case 4:
		fmt.Println("can move 4 spots")
		fallthrough // also go to next case
	case 5:
		fmt.Println("can move 5 spots")
	case 6:
		fmt.Println("can move 6 spots and roll dice again")
	default:
		fmt.Println("what's that!!!")

	}

}
