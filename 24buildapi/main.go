package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware.,helper - file
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API- LCO.in")
	r := mux.NewRouter()

	// seeding

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 199, Author: &Author{Fullname: "Chethan", Website: "go.dev"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "Mern stack", CoursePrice: 299, Author: &Author{Fullname: "Chethan", Website: "go.dev"}})

	//routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	//listen to port

	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers - file
//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One couse")
	w.Header().Set("content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)

	// loop through courses. find matching id that user sending as parameter and return response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create a course")
	w.Header().Set("content-Type", "application/json")

	// what if: body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - { }

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}

	// checking id title is Duplicate
	//params := mux.Vars(r)

	for _, value := range courses {
		if value.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Title already exists")
			return
		}
	}

	// generate unique id, string
	// append course int courses

	rand.NewSource(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update a course")
	w.Header().Set("content-Type", "application/json")
	// first - grab id from req

	params := mux.Vars(r)

	// loop id remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
		}
	}

	//Todo: send a response id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete a course")
	w.Header().Set("contetn-Type", "application/json")

	params := mux.Vars(r)

	//loop , find id , remove(index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode("we have deleted a particular course")
}
