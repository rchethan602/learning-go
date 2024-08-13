package main

import "fmt"

func main() {
	fmt.Println("structs in golang")

	// no inheritance in golang; No super or parent doesn't exist in golang

	Chethan := User{Name: "Chethan", Email: "chethan@go.dev", Status: true, Age: 29}

	fmt.Println(Chethan.Email)

	fmt.Printf("Chethan's details are: %+v", Chethan)

	Sarojamma := User{"Sarojamma", "saro@saro.com", false, 58}
	fmt.Println(Sarojamma.Age)
	object1 := StructName{"Chethan", "cheth@cheth.com", true}
	fmt.Println(object1)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type StructName struct {
	Name      string
	Email     string
	LogggedIn bool
}
