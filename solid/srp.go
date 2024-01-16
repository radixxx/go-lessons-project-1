package main

type Employee struct {
	Name    string
	Salary  float64
	Address string
}

// According to the SRP, each struct should have only one responsibility,
// so in this case, it would be better to split the responsibilities
// If I need to make changes to the salary calculation or address handling,
// I know exactly where to look, without having to wade through a lot of unrelated code.

type EmployeeInfo struct {
	Name   string
	Salary float64
}

type EmployeeAddress struct {
	Address string
}

func (e EmployeeInfo) GetSalary() float64 {
	return e.Salary
}

func (e EmployeeAddress) GetAddress() string {
	return e.Address
}
