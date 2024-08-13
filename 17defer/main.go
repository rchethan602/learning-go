package main

import "fmt"

func main() {
	defer fmt.Println("worlds")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("Three")
	fmt.Println("Hello")
	myDefer()
}

// if there are multiple defere statement in code  last in first out
// defer statement or functions  will be  delayed until execution of surrounding function execution completes
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, "\n")
	}
}
