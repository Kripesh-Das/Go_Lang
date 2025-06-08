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

// Model

type Course struct{
	CourseID string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// fake DB
var courses []Course


// Helper - middleware

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main(){
	r := mux.NewRouter()

	courses = append(courses,Course{CourseID: "2",CourseName: "GenAI" ,CoursePrice: 999 ,
	Author: &Author{Fullname: "Kripesh Das", Website: "ai.dev" }})

	courses = append(courses,Course{CourseID: "4",CourseName: "DevOpsI" ,CoursePrice: 999 ,
	Author: &Author{Fullname: "Kripesh Das", Website: "devops.dev" }})

	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")	
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")	
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")
	
		
	log.Fatal(http.ListenAndServe(":4000", r))
									

}

// Controllers

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to the API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}


func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type","application/json")
	
	// grab id from the request
	params := mux.Vars(r)

	//loop through courses, and find matching id 
	for _,course := range courses{
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Courses Found")
}

func createOneCourse(w http.ResponseWriter, r *http.Request){

	fmt.Println("Create One Course")
	w.Header().Set("Content-Type","application/json")
	// if data is is Empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("PLEASE SEND SOME DATA")
	}

	// what about {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	
	}	
	
	// generate a unique id and convert them into string

	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseID = strconv.Itoa(randSource.Intn(100))
	json.NewEncoder(w).Encode(course)

	


}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-Type","application/json")
	
	// grab id from request
	params := mux.Vars(r)

	//loop ,id ,remove ,add my ID
	for index, course := range courses{

		if course.CourseID == params["id"] {
			courses = append(courses[:index],courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){

	fmt.Println("Delete one Course")
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	// loop, id ,delete
	for index,course := range courses{
		if course.CourseID == params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			break

		}
	}
}