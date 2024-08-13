package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	source := rand.NewSource(time.Now().UnixNano())

	rng := rand.New(source)

	diceNumber := rng.Intn(6) + 1

	fmt.Println("Value of deice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move", diceNumber, "spot")
		fallthrough

	// fallthrough if particular case is true it will continue to execute next available case as well
	case 4:
		fmt.Println("you can move", diceNumber, "spot")
	case 5:
		fmt.Println("you can move", diceNumber, "spot")
	case 6:
		fmt.Println("you can move", diceNumber, "spot")
		fallthrough
	default:
		fmt.Println("what was that?!!")

	}

	/*
		switch value {
		case value1:
			fmt.Println("abc")
		case value2:
			fmt.Println("xyz")
			fallthrough
		default:
			fmt.Println("default")
		}
	*/
}
