package service

import (
	"github.com/eduardomassami/go-crud/data/request"
	"github.com/eduardomassami/go-crud/data/response"
	"github.com/eduardomassami/go-crud/helper"
	"github.com/eduardomassami/go-crud/model"
	"github.com/eduardomassami/go-crud/repository"
	"github.com/go-playground/validator"
)

type EmployeeServiceImpl struct {
	EmployeeRepository repository.EmployeeRepository
	Validate           *validator.Validate
}

func NewEmployeeServiceImpl(employeeRepository repository.EmployeeRepository, validate *validator.Validate) EmployeeService {
	return &EmployeeServiceImpl{
		EmployeeRepository: employeeRepository,
		Validate:           validate,
	}
}

func (t EmployeeServiceImpl) Create(employee request.CreateEmployeeRequest) {
	err := t.Validate.Struct(employee)
	helper.ErrorPanic(err)
	employeeModel := model.Employee{
		Name: employee.Name,
	}
	t.EmployeeRepository.Save(employeeModel)
}

func (t EmployeeServiceImpl) Update(employee request.UpdateEmployeesRequest) {
	employeeData, err := t.EmployeeRepository.FindById(employee.Id)
	helper.ErrorPanic(err)
	employeeData.Name = employee.Name
	t.EmployeeRepository.Update(employeeData)
}

func (t EmployeeServiceImpl) Delete(employeeId int) {
	t.EmployeeRepository.Delete(employeeId)
}

func (t EmployeeServiceImpl) FindById(employeeId int) response.EmployeeResponse {
	employeeData, err := t.EmployeeRepository.FindById(employeeId)
	helper.ErrorPanic(err)

	employeeResponse := response.EmployeeResponse{
		Id:   employeeData.Id,
		Name: employeeData.Name,
	}
	return employeeResponse
}

func (t EmployeeServiceImpl) FindAll() []response.EmployeeResponse {
	result := t.EmployeeRepository.FindAll()

	var Employee []response.EmployeeResponse
	for _, value := range result {
		employee := response.EmployeeResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		Employee = append(Employee, employee)
	}
	return Employee
}
