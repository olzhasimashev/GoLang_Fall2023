package director

import "Assignment_1/employee"

type Director struct {
	position string
	salary   float64
	address  string
}

func (m *Director) GetPosition() string {
	return m.position
}

func (m *Director) SetPosition(p string) {
	m.position = p
}

func (m *Director) GetSalary() float64 {
	return m.salary
}

func (m *Director) SetSalary(s float64) {
	m.salary = s
}

func (m *Director) GetAddress() string {
	return m.address
}

func (m *Director) SetAddress(a string) {
	m.address = a
}

func New() employee.Employee {
	return &Director{position: "Director", salary: 200000, address: "1 Director St"}
}
