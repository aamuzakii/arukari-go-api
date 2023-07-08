package main // should use package main instead of package migrations

import (
	"arukari/initializers"
	"arukari/models"
	"fmt"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("start migrating")
	initializers.DB.AutoMigrate(&models.Employee{}, &models.Company{}, &models.Department{}, &models.Employment{})
}
