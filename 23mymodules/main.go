package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in  Golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	host := r.Header.Get("x-request-downstream")
	//val := strings.Split("127.0.0.1:4000", ":")[0]
	//fmt.Println("val is:", val)
	fmt.Println("host is", r.Host)
	//host := r.Header.Get("X-Forwarded-Host"),
	if host == "" {
		host = r.Header.Get("X-Forwarded-Host")
	}
	if host == "" {
		host = strings.Split(r.Host, ":")[0] // When used with an IP address, Envoy adds port to the host header which we don't expect (case in CI)
	}
	fmt.Println(host)
	w.Write([]byte("<h1>welcome to golang series</h1>"))
}
