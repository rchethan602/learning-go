package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost:80")
	if err != nil {
		log.Fatal(nil)
		os.Exit(1)
	}

	fmt.Println(resp.Body)

	// bs := make([]byte, 99999)

	// resp.Body.Read(bs)

	// println(string(bs))

	io.Copy(os.Stdout, resp.Body)
}
