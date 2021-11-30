package main

import "fmt"

const LoginToken string = "dffasdfgs" // first letter capital means it is public

func main() {
	var username string = "Sayan"
	var isLoggedIn bool = true
	fmt.Println(username)
	fmt.Println(isLoggedIn)

	fmt.Printf("Variable is of type : %T \n", username)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	// implicit type
	var website = "risingleafs.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
