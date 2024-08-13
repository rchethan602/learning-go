package main

import (
	"fmt"
	"math/rand"
	"time"
)

type contactInfo struct {
	email   string
	zipCode int
}

// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo contactInfo
// }

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// jimPointer := &jim // to avoid this unnecessary step we  directly pass the argument with pointer as type see func UpdateFirstName()
	// jimPointer.UpdateFirstName("Jimmy")

	jim.UpdateFirstName("jimmy") // here we are not passing pointer but still it works
	// with go if we define a receiver with a type of pointer to whatever, when we attempt to call this function or
	// we attempt to call this method- go will allow us to either call this function with a pointer or with the type
	jim.Print()

}

func (p person) Print() {
	fmt.Println(p)
}

func (p *person) UpdateFirstName(newFirst string) { // here p can be pointer to  or actual type
	p.firstName = newFirst
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)
	fmt.Println(r.Intn(4))
}

/*
*person - this is type description- it means we're working with pointer to a person
func(pointerToPerson *person) updateFirstName(newFirst string){
	*pointerToPerson (this is an operator -it means we want to manipulate the value the pointe is referencing)
}

turn addres in to value with *address
turn value to address with &value

for slices you don't need to pass pointer , when you modify  a slice , you are not modifying the pointer to head , you are still modi

whe you copy a slice your copy is still referring to a same memory , if you delete a element in copy , in original also it will be removed

slice - pointer to head, Capacity, length

there are similar types like slice in go such as ,slices, maps, channels, pointers, functions -- these are all Reference types  , don't worry about pointers with these

int,float, string, bool, structs - these are all value types , use pointers to change these things in a function
*/
