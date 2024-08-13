package main

import "fmt"

func main() {
	fmt.Println("Lets study pointers")

	// var ptr *int
	// fmt.Println("value of pointer is", ptr)

	myNumber := 23

	var ptr = &myNumber

	var ptr2 = &myNumber

	fmt.Println("value of  actual pointer is", ptr)
	fmt.Println("value of  actual pointer is", *ptr)
	fmt.Println("value of  actual pointer is", *ptr2)

	*ptr = *ptr + 2

	fmt.Println("New Value is:", myNumber)

}

// reference means &
//pointers: passing variables  it creates unwanted set of copies . it creates irregularitties,
// just a direct reference to the memory
// it gives gurantee that operaton will be performed on actual value .

// &variable gives me te memory address of the value this variable is pointing at

//*pointer Gives me the value this memory address is pointing at .
