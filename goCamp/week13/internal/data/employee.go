package data

import "week13/internal/biz"

type EmployeeRepo struct {
	db db
}

func NewEmployeeRepo(db db) biz.EmployeeRepo {
	return &EmployeeRepo{db: db}
}

func (repo *EmployeeRepo) GetById(id string) *biz.Employee {
	return &biz.Employee{}
}

func (repo *EmployeeRepo) GetByName(name string) *biz.Employee {
	return &biz.Employee{}
}
