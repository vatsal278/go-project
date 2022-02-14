package model

type Employee struct {
	Name        string `json:"fullname"`
	Employee_id string `json:"employee_id"`
}

var Employee_slice []Employee
