package developer

import "Assignment_1/employee"

type Developer struct {
	position string
	salary   float64
	address  string
}

func (m *Developer) GetPosition() string {
	return m.position
}

func (m *Developer) SetPosition(p string) {
	m.position = p
}

func (m *Developer) GetSalary() float64 {
	return m.salary
}

func (m *Developer) SetSalary(s float64) {
	m.salary = s
}

func (m *Developer) GetAddress() string {
	return m.address
}

func (m *Developer) SetAddress(a string) {
	m.address = a
}

func New() employee.Employee {
	return &Developer{position: "Developer", salary: 80000, address: "404 Developer St"}
}
