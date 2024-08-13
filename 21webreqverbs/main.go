package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to web verb video-1")
	//PerformGetRequest()
	//PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	var responseStrings strings.Builder

	content, _ := io.ReadAll(response.Body)

	byteCount, _ := responseStrings.Write(content)

	fmt.Println("ByteCount is :", byteCount)
	fmt.Println(responseStrings.String())

	// fmt.Println(string(content))
}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"Country" : "India",
			"course" : "lets go with Golang",
			"platform" : "learncodeonline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"
	//formdata
	data := url.Values{}
	data.Add("firstname", "Chethan")
	data.Add("lastname", "R")
	data.Add("email", "chethan@go.dev")

	response, err := http.PostForm(myurl, data)
	defer response.Body.Close()
	if err != nil {
		panic(nil)
	}

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
