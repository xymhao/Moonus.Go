package biz

type EmployeeRepo interface {
	GetById(id string) *Employee
	GetByName(name string) *Employee
}

type Employee struct {
	options
	IdNumber string
	repo     EmployeeRepo
}

type options struct {
	Name     string
	Age      int
	IdNumber string
}

type Option func(o *options)

func Name(name string) Option {
	return func(o *options) {
		o.Name = name
	}
}

func Age(age int) Option {
	return func(o *options) {
		o.Age = age
	}
}

func NewEmployee(rep EmployeeRepo) *Employee {
	return &Employee{repo: rep}
}

func NewEmployeeOpts(ops ...Option) {
	o := options{}
	for _, opt := range ops {
		opt(&o)
	}
}

func (emp *Employee) Add(name string, age int) bool {
	emp.Name = name
	emp.Age = age
	return true
}
