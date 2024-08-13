package main

import "fmt"

func main() {
	ALPHABETS := "ABCDEFGHIJKLMNO"

	r := []rune(ALPHABETS)

	fmt.Println(r)
	fmt.Println(string(r)[1:5])
}
