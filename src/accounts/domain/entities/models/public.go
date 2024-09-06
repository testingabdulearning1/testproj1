package models

type School struct {
	ID       string `gorm:"primaryKey,id"`
	Name     string `gorm:"name"`
	Location string `gorm:"location"`
	Address  string `gorm:"address"`
	Logo     string `gorm:"logo"`
}

type SuperAdmin struct {
	ID             int    `gorm:"primaryKey,id"`
	Username       string `gorm:"username"`
	HashedPassword string `gorm:"hashedpassword"`
}