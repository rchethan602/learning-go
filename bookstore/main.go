package main

import (
	"bookstore/router"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting book store app")

	http.ListenAndServe(":4000", router.Router())
	fmt.Println("server listing at port :4000")
}
