package config

import (
	"fmt"
	// "practice-api-gin-one/helper"

	"github.com/eduardomassami/go-crud/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = "12345"
	dbName   = "company"
)

func DatabaseConnection() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	helper.ErrorPanic(err)

	return db
}
