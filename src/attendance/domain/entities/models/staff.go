package models

import "time"

type StaffAttendanceEntry struct {
	EntryID    int       `gorm:"primaryKey,entry_id"`
	Date       string    `gorm:"date"`
	StaffID    int       `gorm:"staff_id"`
	Status     string    `gorm:"status"`
	MarkedBy   string    `gorm:"marked_by"`
	MarkedTime time.Time `gorm:"marked_time"`
	IsFinal    bool      `gorm:"is_final"`
}


type StaffLeaveApplication struct {
	ID             int    `gorm:"primaryKey,id"`
	StaffID        int    `gorm:"staff_id"`
	Reason         string `gorm:"reason"`
	LeaveStartDate string `gorm:"leave_start_date"`
	LeaveEndDate   string `gorm:"leave_end_date"`
	ApprovalStatus bool   `gorm:"approval_status"`
}

type StaffAbsenceReason struct {
	ID      int    `gorm:"primaryKey,id"`
	Reason  string `gorm:"reason"`
	EntryId int    `gorm:"entry_id"`
}

