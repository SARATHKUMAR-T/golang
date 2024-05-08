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

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	WebSite  string `json:"website"`
}

// fake db
var courses []Course

// middleware or helper -file

func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to api section of golang")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "reactjs", CoursePrice: 299,
		Author: &Author{FullName: "jonas", WebSite: "jonas.in"},
	})

	courses = append(courses, Course{CourseId: "4", CourseName: "nodejs", CoursePrice: 199,
		Author: &Author{FullName: "hitesh", WebSite: "hitesh.in"},
	})

	// routing
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", allCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":8000", r))
}

// controller -files

// server home controller

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to learn code online</h1>"))
}

// all courses

func allCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("all courses served")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get one course

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-type", "application/json")

	// grap id from request
	params := mux.Vars(r)

	// loop through and find the matching course

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found ")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course ")
	w.Header().Set("Content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("no data inside request")
		return
	}

	// generete unique id,string

	random1 := rand.NewSource((time.Now().UnixNano()))
	random2 := rand.New(random1)
	course.CourseId = strconv.Itoa(random2.Intn(50))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-type", "application/json")

	// grab id from params

	params := mux.Vars(r)

	// loop,id,remove,add

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted successfully")
			break
		}
	}
}
