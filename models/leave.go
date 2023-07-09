package models

import "time"

type LeaveType struct {
	Name        string `gorm:"not null"`
	Description string
	AllowedDays int
}

type LeaveRequest struct {
	EmployeeID  uint      `gorm:"not null"`
	LeaveTypeID uint      `gorm:"not null"`
	StartDate   time.Time `gorm:"not null"`
	EndDate     time.Time `gorm:"not null"`
	Status      string    `gorm:"not null"`
	Reason      string
}
