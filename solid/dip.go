package main

import "fmt"

type Worker struct {
	ID   int
	Name string
}

func (w Worker) GetID() int {
	return w.ID
}

func (w Worker) GetName() string {
	return w.Name
}

type Supervisor struct {
	ID   int
	Name string
}

func (s Supervisor) GetID() int {
	return s.ID
}

func (s Supervisor) GetName() string {
	return s.Name
}

/*
	Dependency Inversion - high-level modules should not depend on low-level modules.
	Instead, both should depend on abstractions.
*/

type Employee interface {
	GetID() int
	GetName() string
}

type Department struct {
	Employees []Employee
}

func (d *Department) AddEmployee(e Employee) {
	d.Employees = append(d.Employees, e)
}

func (d *Department) GetEmployeeNames() (res []string) {
	for _, e := range d.Employees {
		res = append(res, e.GetName())
	}
	return
}

func (d *Department) GetEmployees(id int) Employee {
	for _, e := range d.Employees {
		if e.GetID() == id {
			return e
		}
	}
	return nil
}

func main() {
	dep := &Department{}
	dep.AddEmployee(Worker{ID: 1, Name: "Sam"})
	dep.AddEmployee(Worker{ID: 2, Name: "Christine"})

	fmt.Println(dep.GetEmployeeNames())

	e := dep.GetEmployees(1)
	switch v := e.(type) {
	case Worker:
		fmt.Printf("found worker %+v\n", v)
	case Supervisor:
		fmt.Printf("found supervisor %+v\n", v)
	default:
		fmt.Printf("could not find an employee by id: %d\n", 1)
	}
}

/*
	Department struct depends on an abstraction (the Employee interface),
	rather than on a specific implementation (the Worker or Supervisor struct).
	This makes the code more flexible and easier to maintain, as changes to the
	implementation of workers and supervisors will not affect the Department struct.
*/
