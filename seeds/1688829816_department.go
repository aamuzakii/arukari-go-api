package main

import (
	"arukari/initializers"
	"arukari/models"
)

func main() {

	sales_dept := models.Department{
		Name:        "Sales",
		Description: "Company's spearhead",
	}

	initializers.ConnectToDB()

	initializers.DB.Create(&sales_dept)
}
