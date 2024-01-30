package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Models
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"Author"`
}

type Author struct {
	Fullname string `json:"fulllname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// Middlewares
func (c *Course) IsEmpty() bool {
	return c.CourseName == "" // course id can generate uniquely
	// return c.CourseId == "" && c.CourseName == ""
}

// controllers
func serveHome(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("<h1>Welcome to Go lang learning!</h1>"))
}

func getAllCourses(res http.ResponseWriter, req *http.Request) {
	fmt.Println("INFO - Get All Course Route Hits ---")
	res.Header().Set("Content-Type", "application/json")

	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(courses)
}

func getOneCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("INFO - Get One Course Endpoint Hit ---")
	res.Header().Set("Content-Type", "application/json")
	// get req params
	param := mux.Vars(req)

	// loop through course, find matching id and return single course
	for _, course := range courses {
		if course.CourseId == param["id"] {
			res.WriteHeader(http.StatusOK)
			json.NewEncoder(res).Encode(course)
			return
		}
	}
	res.WriteHeader(http.StatusNotFound)
	json.NewEncoder(res).Encode("No Course found with this id!!") // this is end of function so no need to return
}

func createCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("INFO - Create Course Endpoint Hit ---")
	res.Header().Set("Content-Type", "application/json")

	// Check : body is empty
	if req.Body == nil {
		res.WriteHeader(http.StatusNoContent)
		json.NewEncoder(res).Encode("Please send some JSON data!")
	}

	// Get req body
	var course Course
	json.NewDecoder(req.Body).Decode(&course) // by passing pointer of course it will store decode data into course variable

	// Check : body is {}
	if course.IsEmpty() {
		res.WriteHeader(http.StatusPartialContent)
		json.NewEncoder(res).Encode("Provide required fields coursename")
	}

	// generate unique id
	course.CourseId = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100) + 1)
	// Add into fake db
	courses = append(courses, course)

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(course)
}

func main() {

}
