package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to web verb video -lco")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	performPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	// closing request
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Status code: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body) //content is byte format
	byteCount, _ := responseString.Write(content)

	fmt.Println("byte count is: ", byteCount)
	fmt.Println(string(content))
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"let's go with golang",
			"price":0,
			"platform":"learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func performPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "Sayan")
	data.Add("lastname", "Mondal")
	data.Add("email", "sayan@rl.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
