package main

import (
	"fmt"
)

// jwtToken := 3000 not allowed

const LoginToken string = "scnjdnkwcnwkcn" // Public declaration

func main() {
	var username string = "Chethan"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLogggedIn bool = true
	fmt.Println(isLogggedIn)
	fmt.Printf("variable is of type: %T \n", isLogggedIn)

	var smallVal uint16 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloatVal float64 = 255.45545666
	fmt.Println(smallFloatVal)
	fmt.Printf("variable is of type: %T \n", smallFloatVal)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Varibale is of type: %T \n", anotherVariable)

	//implicit type

	var website = "learncode.in"
	fmt.Println(website)

	// no var style
	// walrus method declaration
	numberOfuser := 30000.002
	fmt.Println(numberOfuser)

	fmt.Println(LoginToken)
	fmt.Printf("Varibale is of type: %T \n", LoginToken)

}
