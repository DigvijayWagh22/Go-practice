package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

func (m *Manager) AddEmployee(e Employee) {
	m.Employees = append(m.Employees, e)
}

func (m *Manager) RemoveEmployee(id int) {
	for idx, e := range m.Employees {
		if e.ID == id {
			m.Employees = append(m.Employees[:idx], m.Employees[idx+1:]...)
			return
		}
	}
}

func (m *Manager) GetAverageSalary() float64 {
	if len(m.Employees) == 0 {
		return 0
	}

	var sum float64
	for _, emp := range m.Employees {
		sum += emp.Salary
	}

	result := sum / float64(len(m.Employees))
	return result
}

func (m *Manager) FindEmployeeByID(id int) *Employee {

	for _, emp := range m.Employees {
		if emp.ID == id {
			return &emp
		}
	}

	return nil
}

func main() {
	manager := Manager{}
	manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 31, Salary: 75000})
	//manager.RemoveEmployee(2)
	averageSalary := manager.GetAverageSalary()
	employee := manager.FindEmployeeByID(2)

	fmt.Printf("Average salary : %f\n", averageSalary)
	if employee != nil {
		fmt.Printf("Employee found: %+v\n", *employee)
	}

}
