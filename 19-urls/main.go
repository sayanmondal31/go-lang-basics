package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=54545sdfd"

func main() {
	fmt.Println("welcome to urls")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of quesry params are: %T\n", qparams)

	// extracting details from url
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["payment"])

	//for index,val
	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
