package models

type TermExam struct {
	ID       int    `gorm:"primaryKey,id"`
	Type     string `gorm:"type"`
	Name     string `gorm:"name"`
	Standard int    `gorm:"standard"`
}