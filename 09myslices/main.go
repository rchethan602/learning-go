package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"apple", "Tomato", "peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fmt.Println("len of fruitlist is", len(fruitList))

	fruitList = append(fruitList, "Mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 04
	highscores[2] = 14
	highscores[3] = 23

	//highscores[4] = 24
	//highscores[5] = 34
	fmt.Println(len(highscores))
	highscores = append(highscores, 555, 666)
	fmt.Println(len(highscores))
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println(highscores)

	// how to remove values from slices based on index

	var courses = []string{"react", "js", "swift", "python", "ruby"}

	var slice1 = []int{1, 2, 3, 4, 5, 6, 7}
	var s0 = &slice1
	fmt.Printf("address is", s0)
	//fmt.Println(&slice1)

	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	slice1 = append(slice1, slice1[0:4]...)
	// s1 = &slice1
	fmt.Printf("address is %v", slice1)

}

// end index not included when you slince
