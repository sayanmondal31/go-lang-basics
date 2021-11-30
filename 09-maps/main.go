package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript" // must use double quotes
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages ", languages)
	fmt.Println("Js shorts for ", languages["JS"])

	// deleting value from maps
	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("for key v, value is %v\n", value)
	}
}
