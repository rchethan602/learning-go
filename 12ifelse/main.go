package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "watch out!!"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// you can also declare varibales in the condition and you can use it .

	if num := 3; num < 10 {
		fmt.Println("Num us less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil {

	// }
}
