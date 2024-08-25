package main

import (
	"net/http"
	"time"

	"github.com/eduardomassami/go-crud/config"
	"github.com/eduardomassami/go-crud/controller"
	"github.com/eduardomassami/go-crud/helper"
	"github.com/eduardomassami/go-crud/model"
	"github.com/eduardomassami/go-crud/repository"
	"github.com/eduardomassami/go-crud/router"
	"github.com/eduardomassami/go-crud/service"
	"github.com/go-playground/validator"
)

func main() {
	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("employees").AutoMigrate(&model.Employee{})

	//Init Repository
	employeeRepository := repository.NewEmployeeRepositoryImpl(db)

	//Init Service
	employeeService := service.NewEmployeeServiceImpl(employeeRepository, validate)

	//Init controller
	employeeController := controller.NewEmployeeController(employeeService)

	//Router
	routes := router.NewRouter(employeeController)

	server := &http.Server{
		Addr:           ":8081",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
