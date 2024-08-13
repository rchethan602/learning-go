package main

import "fmt"

func main() {
	fmt.Println("methods in golang")

	// no inheritance in golang; No super or parent doesn't exist in golang

	Chethan := User{Name: "Chethan", Email: "chethan@go.dev", Status: true, Age: 29}

	fmt.Println(Chethan.Email)

	fmt.Printf("Chethan's details are: %+v\n", Chethan)
	Chethan.GetStatus()
	Chethan.NewMail()
	fmt.Printf("Chethan's details are: %+v\n", Chethan)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New email of the user is: ", u.Email)
}
