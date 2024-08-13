package main

import (
	"bufio"
	"fmt"
	"os"

	"main2.go/alter"
)

func main() {
	welcome := "welcome to user p"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
	alter.Main1()
}
