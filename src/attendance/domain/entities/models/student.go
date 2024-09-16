package models

import "time"

type StudentAttendanceInstance struct {
	ID      int    `gorm:"primaryKey,id"`
	ClassID int    `gorm:"class_id"`
	Date    string `gorm:"date"`
	TakenBy int    `gorm:"taken_by"`
}

type StudentAttendanceEntry struct {
	EntryID    int    `gorm:"primaryKey,entry_id"`
	InstanceID int    `gorm:"instance_id"`
	StudentID  int    `gorm:"student_id"`
	Status     string `gorm:"status"`
}

type StudentAttendanceEdits struct {
	InstanceID    int       `gorm:"instance_d"`
	StudentID     int       `gorm:"student_id"`
	InitialStatus string    `gorm:"initial_status"`
	FinalStatus   string    `gorm:"final_status"`
	EditedStaffID int       `gorm:"edited_staff_id"`
	EdittedTime   time.Time `gorm:"edited_time"`
}

// type StudentAbsenceReason struct {	//For future
// 	StudentID     int    `gorm:"student_id"`
// 	Reason        string `gorm:"reason"`
// 	LeaveStartDate string `gorm:"leave_start_date"`
// 	LeaveEndDate   string `gorm:"leave_end_date"`
// }
// type StudentLeaveApplication struct { //For future
// 	ID             int    `gorm:"primaryKey,id"`
// 	StudentID      int    `gorm:"student_id"`
// 	Reason         string `gorm:"reason"`
// 	LeaveStartDate string `gorm:"leave_start_date"`
// 	LeaveEndDate   string `gorm:"leave_end_date"`
// 	ApprovalStatus bool   `gorm:"approval_status"`
// }
