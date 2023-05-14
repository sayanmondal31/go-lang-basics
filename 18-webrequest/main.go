package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response type is : %T\n", response)

	defer response.Body.Close()

	// read response with ioutils
	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println(content)
}
