package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://localhost"

const url1 = "https://lco.dev"

func main() {
	fmt.Println("web requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of Type: %T\n", response)
	defer response.Body.Close()

	databytes, err := io.ReadAll((response.Body))

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)

}
