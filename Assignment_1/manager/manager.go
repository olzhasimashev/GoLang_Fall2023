package manager

import "Assignment_1/employee"

type Manager struct {
	position string
	salary   float64
	address  string
}

func (m *Manager) GetPosition() string {
	return m.position
}

func (m *Manager) SetPosition(p string) {
	m.position = p
}

func (m *Manager) GetSalary() float64 {
	return m.salary
}

func (m *Manager) SetSalary(s float64) {
	m.salary = s
}

func (m *Manager) GetAddress() string {
	return m.address
}

func (m *Manager) SetAddress(a string) {
	m.address = a
}

func New() employee.Employee {
	return &Manager{position: "Manager", salary: 100000, address: "123 Manager St"}
}
