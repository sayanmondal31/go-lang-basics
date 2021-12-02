package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("file")

	content := "This need to go in a file - rl.com"

	file, err := os.Create("./myrlgofile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./myrlgofile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
