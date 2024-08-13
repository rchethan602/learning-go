package main

import "fmt"

func main() {
	fmt.Println("welocme to loops in Golang")

	days := []string{"Monday", "Tuesday", "wednesday", "friday", "saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rougueValue := 2

	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
			//goto cannot jump to a label that is outside the current function.
		}
		if rougueValue == 5 {
			fmt.Println(rougueValue)
			rougueValue++
			fmt.Println(rougueValue)
			continue
			//The continue statement is used within loops to skip the current iteration and proceed to the next iteration
			// break
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("jumping at learningcodeonline.in")
}
