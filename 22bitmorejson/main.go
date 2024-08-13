package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to Json video")

	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcocourses := []course{
		{"ReactJs Bootcamp", 100, "youtube", "pass123", []string{"web-dev", "js"}},
		{"angular Bootcamp", 200, "youtube", "pass123", []string{"web-dev", "js"}},
		{"vennila Bootcamp", 300, "youtube", "pass123", []string{"web-dev", "js"}},
		{"nodejs Bootcamp", 400, "youtube", "pass123", nil},
	}
	// package this data as JSON code

	finalJson, err := json.MarshalIndent(lcocourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

	//fmt.Println(string(finalJson)), both are same

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs Bootcamp",
		"Price": 100,
		"website": "youtube",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		//fmt.Printf("%#v", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// some cases where you just want to add data to key value pair

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and TYpe is: %T\n", k, v, v)
	}

}
