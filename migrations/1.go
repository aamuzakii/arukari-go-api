package main // should use package main instead of package migrations

import (
	"arukari/initializers"
	"arukari/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Employee{}, &models.Company{}, &models.Department{}, &models.Employment{})
}
