package main

import "fmt"

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

func embedStruct() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)
}
