package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// conversion in golang
func main() {
	fmt.Println("welcome to pur pizza app")
	fmt.Println("please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("error occured")
	} else {

		fmt.Println("added 1 to your rating: ", numRating+1)
	}

}
