package hr

import "Assignment_1/employee"

type Hr struct {
	position string
	salary   float64
	address  string
}

func (m *Hr) GetPosition() string {
	return m.position
}

func (m *Hr) SetPosition(p string) {
	m.position = p
}

func (m *Hr) GetSalary() float64 {
	return m.salary
}

func (m *Hr) SetSalary(s float64) {
	m.salary = s
}

func (m *Hr) GetAddress() string {
	return m.address
}

func (m *Hr) SetAddress(a string) {
	m.address = a
}

func New() employee.Employee {
	return &Hr{position: "Hr", salary: 50000, address: "41 Hr St"}
}
