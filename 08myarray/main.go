package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "jackfruit"

	fmt.Println("the lenght of fruitlist is", len(fruitList))

	var grades = [...]int{10, 20, 30}
	fmt.Println("the length of grades is", len(grades))

}
