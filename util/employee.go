package util

import (
	"arukari/initializers"
	"arukari/models"
)

func GetEmployee(email string) (models.Employee, error) {
	var employee models.Employee

	tx := initializers.DB.Where("email = ?", email).First(&employee)

	return employee, tx.Error
}
