package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name string
}

type Employee struct {
	gorm.Model
	Name       string
	Department Department `gorm:"foreignKey:DepartmentID"`
	Position   string
}
