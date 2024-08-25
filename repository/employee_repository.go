package repository

import "github.com/eduardomassami/go-crud/model"

type EmployeeRepository interface {
	Save(employees model.Employee)
	Update(employees model.Employee)
	Delete(employeesId int)
	FindById(employeesId int) (employees model.Employee, err error)
	FindAll() []model.Employee
}
