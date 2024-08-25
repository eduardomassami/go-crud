package repository

import (
	"errors"

	"github.com/eduardomassami/go-crud/data/request"
	"github.com/eduardomassami/go-crud/helper"
	"github.com/eduardomassami/go-crud/model"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	Db *gorm.DB
}

func NewEmployeeRepositoryImpl(Db *gorm.DB) EmployeeRepository {
	return &EmployeeRepositoryImpl{Db: Db}
}

func (t EmployeeRepositoryImpl) Save(Employee model.Employee) {
	result := t.Db.Create(&Employee)
	helper.ErrorPanic(result.Error)

}

func (t EmployeeRepositoryImpl) Update(Employee model.Employee) {
	var updateEmployee = request.UpdateEmployeesRequest{
		Id:   Employee.Id,
		Name: Employee.Name,
	}
	result := t.Db.Model(&Employee).Updates(updateEmployee)
	helper.ErrorPanic(result.Error)
}

func (t EmployeeRepositoryImpl) Delete(EmployeeId int) {
	var Employee model.Employee
	result := t.Db.Where("id = ?", EmployeeId).Delete(&Employee)
	helper.ErrorPanic(result.Error)
}

func (t EmployeeRepositoryImpl) FindById(EmployeeId int) (model.Employee, error) {
	var employee model.Employee
	result := t.Db.Find(&employee, EmployeeId)
	if result != nil {
		return employee, nil
	} else {
		return employee, errors.New("employee is not found")
	}
}

func (t EmployeeRepositoryImpl) FindAll() []model.Employee {
	var Employee []model.Employee
	results := t.Db.Find(&Employee)
	helper.ErrorPanic(results.Error)
	return Employee
}
