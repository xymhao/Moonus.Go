package biz

type EmployeeRepo interface {
	GetById(id string) *Employee
	GetByName(name string) *Employee
}

type Employee struct {
	Name     string
	Age      int
	IdNumber string
	repo     EmployeeRepo
}

func NewEmployee(rep EmployeeRepo) *Employee {
	return &Employee{repo: rep}
}

func (emp *Employee) Add(name string, age int) bool {
	emp.Name = name
	emp.Age = age
	return true
}
