package models

type Subject struct {
	ID   string `gorm:"primaryKey,id"`
	Name string `gorm:"name"`
}

type Books struct {
	ID         string `gorm:"primaryKey,id"`
	Standard   string `gorm:"standard"`
	SyllabusID int    `gorm:"syllabusId"`
	Name       string `gorm:"name"`
	Link       string `gorm:"link"`
	Medium     string `gorm:"medium"`
}

type Syllabus struct {
	ID   int    `gorm:"primaryKey,id"`
	Name string `gorm:"name"`
}

type Standards struct {
	ID       int    `gorm:"primaryKey,id"`
	Standard string `gorm:"standard"`
}
