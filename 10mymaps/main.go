package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["sh"] = "shellscriptt"
	languages["py"] = "python"

	fmt.Println("map of languaages:", languages)
	fmt.Println("JS shorts for:", languages["js"])
	//deleting

	delete(languages, "py")
	fmt.Println("map of languaages:", languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("for Key %v, value is %v\n", key, value)
	}

	//persons := make(map[string]int)

	persons := map[string]int{"che": 1, "tha": 2, "nr": 3}
	fmt.Println(len(persons))

	for key, value := range persons {
		fmt.Println(key, *&value)
	}

}
