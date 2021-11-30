package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")

	var fruitList = []string{"Apple", "Peach", "Dragonfruit"} // initializing slices

	fmt.Printf("Type of fruitlist is %T\n", fruitList) // shows the type of fruitlist

	// adding new fruits to fruitlist
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// another way of initializing slice with allocation of memory

	highScores := make([]int, 4)

	highScores[0] = 230
	highScores[1] = 303
	highScores[2] = 950
	highScores[3] = 238
	// highScores[3] = 238

	highScores = append(highScores, 555, 352, 450) // even we allocating memory with 4 but while using append method
	// go reinitializing memory

	fmt.Println(highScores)
	fmt.Println("is my slice sorted? ", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println("is my slice sorted? ", sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"Reactjs", "Javascript", "Swift", "Python", "Ruby"}

	fmt.Println(courses)

	// remove course
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("new modified courses ", courses)

}
