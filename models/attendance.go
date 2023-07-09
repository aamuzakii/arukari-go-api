package models

import "time"

// Attendance model
type Attendance struct {
	EmployeeID   uint      `gorm:"not null"`
	Date         time.Time `gorm:"not null"`
	Status       string    `gorm:"not null"`
	ClockInTime  time.Time
	ClockOutTime time.Time
	TotalHours   float64
}

// AttendanceLog model
type AttendanceLog struct {
	EmployeeID uint      `gorm:"not null"`
	Date       time.Time `gorm:"not null"`
	LogType    string    `gorm:"not null"`
	LogTime    time.Time `gorm:"not null"`
	Remarks    string
}
