package service

import (
	"context"
	"strconv"
	"week13/api"
	"week13/api/user"
	"week13/internal/biz"
)

type EmpService struct {
	user.UnimplementedEmployeeServer

	repo     biz.EmployeeRepo
	employee *biz.Employee
}

func NewEmpService(repo biz.EmployeeRepo, employee *biz.Employee) api.EmployeeServer {
	return &EmpService{repo: repo, employee: employee}
}
func NewEmpService2(repo biz.EmployeeRepo, employee *biz.Employee) *EmpService {
	return &EmpService{repo: repo, employee: employee}
}

func (svc EmpService) Add2(emp map[string]string) error {
	age, err := strconv.Atoi(emp["age"])
	if err != nil {
		return err
	}
	svc.employee.Add(emp["name"], age)
	return nil
}
func (svc *EmpService) Add(ctx context.Context, in *user.AddRequest) (*user.AddReply, error) {
	return &user.AddReply{Message: "hello"}, nil
}

func Create() {
	biz.NewEmployeeOpts(biz.Name("moonus"), biz.Age(27))
}
