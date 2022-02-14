package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vatsal278/go-project/model"
	"github.com/gorilla/mux"
)

func StartPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API with all CRUD operations (Employee system)</h1>"))
}

func DisplayEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, employee := range model.Employee_slice {
		if employee.Employee_id == param["id"] {
			json.NewEncoder(w).Encode(employee)
			return
		}
	}
	json.NewEncoder(w).Encode("No Employee with given employee id")
	return
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee model.Employee
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&employee)
	if employee.Employee_id == "" {
		json.NewEncoder(w).Encode("employee id cannot be empty")
		return
	}
	for _, singlemeployee := range model.Employee_slice {
		if employee.Employee_id == singlemeployee.Employee_id {
			json.NewEncoder(w).Encode("Employee id already exists")
			return
		}
	}
	model.Employee_slice = append(model.Employee_slice, employee)
	json.NewEncoder(w).Encode("Successfully added in DB")
	json.NewEncoder(w).Encode(employee)
	return
}
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	for index, employee := range model.Employee_slice {
		if employee.Employee_id == param["id"] {
			model.Employee_slice = append(model.Employee_slice[:index], model.Employee_slice[index+1:]...)
			json.NewEncoder(w).Encode("Employee Deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("No such employee id present")
}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	for index, employee := range model.Employee_slice {
		if employee.Employee_id == param["id"] {
			model.Employee_slice = append(model.Employee_slice[:index], model.Employee_slice[index+1:]...)
			fmt.Println("Employee found to be updated")
			fmt.Println(employee)
			x := employee.Employee_id
			var updatedemployee model.Employee
			json.NewDecoder(r.Body).Decode(&updatedemployee)
			if updatedemployee.Name == "" || updatedemployee.Position == "" || updatedemployee.Salary == 0{
				json.NewEncoder(w).Encode("No field can be nil except Employee id")
				return
			}
			updatedemployee.Employee_id = x
			model.Employee_slice = append(model.Employee_slice, updatedemployee)
			json.NewEncoder(w).Encode("Employee has been updated successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No employee with the given id is present in the database to be updated")

}
