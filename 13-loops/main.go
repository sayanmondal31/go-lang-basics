package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in go line")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "friday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("index is v and value is %v\n", day)
	// }

	// like while loop
	dummyValue := 1
	for dummyValue <= 10 {

		if dummyValue == 2 {
			goto rl
		}
		if dummyValue == 5 {
			dummyValue++
			continue
		}

		if dummyValue == 8 {
			break
		}
		fmt.Println("value is: ", dummyValue)
		dummyValue++
	}

	// using it fot goto
rl:
	fmt.Println("Jumping at risingleafs.com")

}
