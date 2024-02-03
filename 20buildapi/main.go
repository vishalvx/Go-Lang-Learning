package main

import (
	"encoding/json"
	"log"
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
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fulllname"`
	Website  string `json:"website"`
}

// type CourseHandler struct {
// 	http.Handler
// }

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
	log.Print("INFO - Get All Course Route Hits ---")
	res.Header().Set("Content-Type", "application/json")

	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(courses)
}

func getOneCourse(res http.ResponseWriter, req *http.Request) {
	log.Print("INFO - Get One Course Endpoint Hit ---")
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
	log.Print("INFO - Create Course Endpoint Hit ---")
	res.Header().Set("Content-Type", "application/json")

	// Check : body is empty
	if req.Body == nil {
		res.WriteHeader(http.StatusNoContent)
		json.NewEncoder(res).Encode("Please send some JSON data!")
	}

	// Get req body
	var course Course
	json.NewDecoder(req.Body).Decode(&course) // by passing pointer of course it will store decode data into course variable

	// Check : body is {} or course name is not provide
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

func updateOneCourse(res http.ResponseWriter, req *http.Request) {
	log.Print("INFO - Update Course Endpoint Hit ---")
	res.Header().Set("Content-Type", "application/json")

	// Get Id from params
	params := mux.Vars(req)

	//loop through course find course
	// add remove that found course and add update data with
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course

			// Case : body empty
			if err := json.NewDecoder(req.Body).Decode(&course); err != nil {
				res.WriteHeader(http.StatusNoContent)
				json.NewEncoder(res).Encode("Please send some JSON data!")
				return
			}
			course.CourseId = params["id"]

			// add new course with same id
			courses = append(courses, course)

			res.WriteHeader(http.StatusOK)
			json.NewEncoder(res).Encode(course)
			return
		}
	}

	// Case : id not found
	res.WriteHeader(http.StatusNotFound)
}

func deleteOneCourse(res http.ResponseWriter, req *http.Request) {
	log.Print("INFO - Delete One Course Endpoint Hit ---")
	res.Header().Set("Content-Type", "application/json")

	// Grab id of course
	params := mux.Vars(req)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			res.WriteHeader(http.StatusOK)
			json.NewEncoder(res).Encode("Course delete successfully!!")
			return
		}
	}
	// Case : id not found
	res.WriteHeader(http.StatusNotFound)
}

func main() {
	log.Print("SERVER --- Starting...")

	// Creating Some Course
	courses = append(courses, Course{CourseId: "1", CourseName: "Learn Go", CoursePrice: 200, Author: &Author{Fullname: "Vishal Prajapati", Website: "https://vishalvx.dev"}}, Course{CourseId: "2", CourseName: "Learn Mongo DB", CoursePrice: 250, Author: &Author{Fullname: "Vishal Prajapati", Website: "https://vishalvx.dev"}})

	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/courses", createCourse).Methods("POST")
	router.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")
	// subRouter := router.NewRoute().Subrouter()
	// subRouter.HandleFunc("/course", CourseHandler)

	log.Print("Listening on http://localhost:3000 ...")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal("Fail to Start Server!!")
	}

}
