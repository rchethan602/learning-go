package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")
	content := "This needs to go in a file - learncodeOnline.in"

	file, err := os.Create("./mylogfile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is:", length)
	defer file.Close()

	readFile("./mylogfile.txt")
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println(string(data))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
