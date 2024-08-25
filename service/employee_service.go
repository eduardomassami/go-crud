package service

import (
	"github.com/eduardomassami/go-crud/data/request"
	"github.com/eduardomassami/go-crud/data/response"
)

type EmployeeService interface {
	Create(Employees request.CreateEmployeeRequest)
	Update(Employees request.UpdateEmployeesRequest)
	Delete(EmployeesId int)
	FindById(EmployeesId int) response.EmployeeResponse
	FindAll() []response.EmployeeResponse
}
