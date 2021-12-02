package main

import "fmt"

func main() {
	fmt.Println("struct in go lang")

	// no inheritance in go lang ; no super or parent
	sayan := User{"Sayan", "sayan@rl.com", true, 16}

	fmt.Println(sayan)
	fmt.Printf("Sayan details : %+v \n ", sayan)
	fmt.Printf("Name is %v and Email is %v \n", sayan.Name, sayan.Email)
	sayan.GetStatus()
	sayan.NewMail()
}

// make capital while creating structure
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// if i want to export then G of getstatus will will be caps
func (u User) GetStatus() { // here is u is the copy of User Structure
	fmt.Println("is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@test.com"
	fmt.Println("Email of this user is: ", u.Email)
}
