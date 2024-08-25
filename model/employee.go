package model

type Employee struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(50)"`
}
