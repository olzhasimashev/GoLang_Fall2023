package designer

import "Assignment_1/employee"

type Designer struct {
	position string
	salary   float64
	address  string
}

func (m *Designer) GetPosition() string {
	return m.position
}

func (m *Designer) SetPosition(p string) {
	m.position = p
}

func (m *Designer) GetSalary() float64 {
	return m.salary
}

func (m *Designer) SetSalary(s float64) {
	m.salary = s
}

func (m *Designer) GetAddress() string {
	return m.address
}

func (m *Designer) SetAddress(a string) {
	m.address = a
}

func New() employee.Employee {
	return &Designer{position: "Designer", salary: 90000, address: "101 Designer St"}
}
