package controller

import (
	"net/http"
	"strconv"

	"github.com/eduardomassami/go-crud/data/request"
	"github.com/eduardomassami/go-crud/data/response"
	"github.com/eduardomassami/go-crud/helper"
	"github.com/eduardomassami/go-crud/service"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	Employeeservice service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{Employeeservice: service}
}

func (controller *EmployeeController) Create(ctx *gin.Context) {
	createEmployeeRequest := request.CreateEmployeeRequest{}
	err := ctx.ShouldBindJSON(&createEmployeeRequest)
	helper.ErrorPanic(err)

	controller.Employeeservice.Create(createEmployeeRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *EmployeeController) Update(ctx *gin.Context) {
	updateEmployeeRequest := request.UpdateEmployeesRequest{}
	err := ctx.ShouldBindJSON(&updateEmployeeRequest)
	helper.ErrorPanic(err)

	employeeId := ctx.Param("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.ErrorPanic(err)

	updateEmployeeRequest.Id = id

	controller.Employeeservice.Update(updateEmployeeRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *EmployeeController) Delete(ctx *gin.Context) {
	employeeId := ctx.Param("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.ErrorPanic(err)
	controller.Employeeservice.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *EmployeeController) FindById(ctx *gin.Context) {
	employeeId := ctx.Param("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.ErrorPanic(err)

	employeeResponse := controller.Employeeservice.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   employeeResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *EmployeeController) FindAll(ctx *gin.Context) {
	employeeResponse := controller.Employeeservice.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   employeeResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
