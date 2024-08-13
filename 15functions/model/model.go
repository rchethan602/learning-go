package model

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Score     int
}

func (p Person) Print() {
	fmt.Println(p.FirstName)
}
