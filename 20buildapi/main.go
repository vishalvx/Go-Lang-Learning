package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	return c.CourseId == "" && c.CourseName == ""
}

// controllers
func serveHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Welcome to Go lang learning!</h1>"))
}

func getAllCourses(res http.ResponseWriter, req *http.Request) {
	fmt.Println("INFO - Get All Course Route Hits ---")
	res.Header().Set("Content-Type", "application/json")
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
			json.NewEncoder(res).Encode(course)
			return
		}
	}
	json.NewEncoder(res).Encode("No Course found with this id!!")
	return
}

func main() {

}
