package main // should use package main instead of package migrations

import (
	"arukari/initializers"
	"arukari/models"
	"log"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	log.Println("start migrating")
	initializers.DB.AutoMigrate(
		&models.Employee{},
		&models.Company{},
		&models.Department{},
		&models.Employment{},
		&models.Attendance{},
		&models.AttendanceLog{},
		&models.LeaveRequest{},
		&models.LeaveType{},
	)
}
