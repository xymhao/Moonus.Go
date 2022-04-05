package service

import (
	"Moonus.Go/goCamp/week4/api"
	"Moonus.Go/goCamp/week4/internal/biz"
	"strconv"
)

type EmpService struct {
	repo     biz.EmployeeRepo
	employee *biz.Employee
}

func NewEmpService(repo biz.EmployeeRepo, employee *biz.Employee) api.EmployeeServer {
	return &EmpService{repo: repo, employee: employee}
}

func (svc EmpService) Add(emp map[string]string) error {
	age, err := strconv.Atoi(emp["age"])
	if err != nil {
		return err
	}
	svc.employee.Add(emp["name"], age)
	return nil
}

func Create() {
	biz.NewEmployeeOpts(biz.Name("moonus"), biz.Age(27))
}
