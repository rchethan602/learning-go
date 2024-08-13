package main

import (
	"fmt"
	"functions/model"
)

func main() {
	fmt.Println("welcome to  fucntions in golang")

	greeter()

	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	proRes, str := proAdder(2, 3, 4, 5, 6)
	fmt.Println("Proresult is", proRes, str)

	chethan := model.Person{FirstName: "chethan", LastName: "R", Score: 20}

	sindhu := model.Person{FirstName: "Sindhu", LastName: "R", Score: 35}

	chethan.Print()

	sindhu.Print()

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result function"
}
func greeter() {
	fmt.Println("Namastey from golang")
}

//signatures  what kind of values are expected and what kind of values are returned
