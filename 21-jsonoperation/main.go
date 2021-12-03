package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"coursename"` // this are used as alias
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` // if there is null in tags omitempty don't show that field
}

func main() {
	fmt.Println("json operation")
	// encodeJson()
	decodeJson()
}

func encodeJson() {

	coursesList := []Course{
		{"React", 299, "learnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"React native", 199, "learnCodeOnline.in", "affbc123", []string{"mobile-dev", "js"}},
		{"Flutter", 399, "learnCodeOnline.in", "abc123ghg", nil},
	} //courses slices

	// package this data into json

	finalJson, err := json.MarshalIndent(coursesList, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

// consume json data
func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React",
		"Price": 299,
		"website": "learnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse Course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("json Was valid ")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	// another usecases
	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
