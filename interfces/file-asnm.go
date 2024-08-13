package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := os.Args

	fmt.Println(fileName[1])
	readFile(fileName[1])
}

func readFile(f string) {

	content, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
