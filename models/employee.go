package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name        string
	Description string
}

type Employee struct {
	gorm.Model
	Name         string
	Email        string `gorm:"unique"`
	Password     string
	KTPNumber    uint
	DepartmentID uint
	Department   Department `gorm:"foreignKey:DepartmentID"`
	Position     string
	Employment   Employment
}

type Employment struct {
	gorm.Model
	CompanyID  uint
	Company    Company `gorm:"foreignKey:CompanyID"`
	EmployeeID uint
}

type Company struct {
	gorm.Model
	Name string
}
