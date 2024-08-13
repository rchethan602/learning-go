package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointers

var mut sync.Mutex //pointer

var signals = []string{"test"}

func main() {
	// go greeter("Hello")
	// greeter("world")

	websiteList := []string{
		"https://lco.dev",
		"https://google.co.in",
		"https://github.com",
		"http://localhost:80",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1) // how may number of calls happening
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(s)
// 	}
// }

// THreads are managed by OS

// Go routines are manged by go run time, go will go ahead and fire up threads

func getStatusCode(endpoint string) {

	defer wg.Done() // gets executed at the end
	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOps in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		fmt.Printf("%d status codefor %s \n", result.StatusCode, endpoint)
		mut.Unlock()
	}
}

// wait groups are used to achive paralleslism effectiviely
