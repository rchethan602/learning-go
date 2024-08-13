package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gjsnknkmsk"

func main() {
	fmt.Println("welcome to handline URLs in golang")
	fmt.Println(myurl)

	// parsing

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	//fmt.Println(result.Host, "\n", result.Path, "\n", result.Port(), "\n")
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf(" type of  qparams:%T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "http",
		Host:    "localhost",
		Path:    "/golang",
		RawPath: "user=chethan",
	}

	anotherURL := partsOfUrl.String()

	fmt.Println(anotherURL)

}
