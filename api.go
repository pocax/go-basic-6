package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Employee struct {
	Id       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{Id: 1, Name: "John", Age: 30, Division: "IT"},
	{Id: 2, Name: "Mary", Age: 25, Division: "HR"},
	{Id: 3, Name: "Mike", Age: 27, Division: "IT"},
}

var PORT = ":8181"

func main() {
	http.HandleFunc("/employees", GetEmployees)

	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Server started at port " + PORT)
	http.ListenAndServe(PORT, nil)
}

// func GetEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "GET" {
// 		json.NewEncoder(w).Encode(employees)
// 		return
// 	}

// 	http.Error(w, "Invalid request method", http.StatusBadRequest)
// }

//use html
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, employees)
		return
	}
	http.Error(w, "Invalid request method", http.StatusBadRequest)

}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			Id:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)
		json.NewEncoder(w).Encode(employees)
		return
	}
	http.Error(w, "Invalid request method", http.StatusBadRequest)

}
