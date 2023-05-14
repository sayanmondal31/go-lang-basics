package main

import "fmt"

func main() {
	fmt.Println("struct in go lang")

	// no inheritance in go lang ; no super or parent
	sayan := User{"Sayan", "sayan@rl.com", true, 16}

	fmt.Println(sayan)
	fmt.Printf("Sayan details : %+v \n ", sayan)
	fmt.Printf("Name is %v and Email is %v ", sayan.Name, sayan.Email)
}

// make capital while creating structure ,also can be accessed by anybody
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
