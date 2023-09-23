// main.go

package main

import (
	"Assignment_1/designer"
	"Assignment_1/developer"
	"Assignment_1/director"
	"Assignment_1/employee"
	"Assignment_1/hr"
	"Assignment_1/manager"
	"fmt"
)

func printDetails(e employee.Employee) {
	fmt.Println("Position:", e.GetPosition())
	fmt.Println("Salary:", e.GetSalary())
	fmt.Println("Address:", e.GetAddress())
}

func main() {
	var employees []employee.Employee

	employees = append(employees, manager.New())
	employees = append(employees, developer.New())
	employees = append(employees, hr.New())
	employees = append(employees, director.New())
	employees = append(employees, designer.New())

	for _, e := range employees {
		printDetails(e)
		fmt.Println("------")
	}
}
