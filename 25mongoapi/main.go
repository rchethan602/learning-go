package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rchethan602/mongoapi/router"
)

func main() {
	fmt.Println("MongoDb API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port :4000")
}
