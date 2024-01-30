package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func main() {

}
