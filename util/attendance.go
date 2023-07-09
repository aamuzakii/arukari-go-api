package util

import (
	"arukari/initializers"
	"arukari/models"
	"time"
)

func CreateAttendance(id uint) error {

	now := time.Now()

	attendance := models.Attendance{
		EmployeeID:   id,
		Date:         now,
		Status:       "ok",
		ClockInTime:  now,
		ClockOutTime: now,
		TotalHours:   12,
	}

	log := models.AttendanceLog{
		EmployeeID: id,
		Date:       now,
		LogType:    "in",
		LogTime:    now,
		Remarks:    "",
	}

	tx := initializers.DB.Create(&attendance)
	logTx := initializers.DB.Create(&log)

	if tx.Error != nil {
		return tx.Error
	}
	if logTx.Error != nil {
		return logTx.Error
	}
	return nil

}
