package models

type Student struct {
	ID             int    `gorm:"primaryKey,id"`
	Password       string `gorm:"password"`
	IdCardNumber   string `gorm:"idcardnumber"`
	Name           string `gorm:"name"`
	ClassId        string `gorm:"classId"`
	Gender         string `gorm:"gender"`
	Dob            string `gorm:"dob"`
	Mobile         string `gorm:"mobile"`
	ParentsInfoId  int    `gorm:"parent1Id"`
	Address        string `gorm:"address"`
	BloodGroup     string `gorm:"bloodGroup"`
	SecondLanguage string `gorm:"second_language"`
	ProfilePicLink string `gorm:"profile_pic_Link"`
	Remarks        string `gorm:"remarks"`
	Religion       string `gorm:"religion"`
	Caste          string `gorm:"caste"`
	Pwd            bool   `gorm:"PWD"`
	Nationality    string `gorm:"nationality"`
}

type Staff struct {
	ID               int    `gorm:"primaryKey,id"`
	Name             string `gorm:"name"`
	HashedPassword   string `gorm:"hashedpassword"`
	Address          string `gorm:"address"`
	Mobile           string `gorm:"mobile"`
	IsTeachingStaff  bool   `gorm:"isTeachingStaff"`
	AdditionalRole   string `gorm:"additionalRole"`
}

type Admin struct {
	ID             int    `gorm:"primaryKey,id"`
	HashedPassword string `gorm:"hashedpassword"`
	Name           string `gorm:"name"`
	Mobile         string `gorm:"mobile"`
	Address        string `gorm:"address"`
}

type Parent struct {
	ID           int    `gorm:"primaryKey,id"`
	MotherName   string `gorm:"motherName"`
	MotherMobile string `gorm:"motherMobile"`
	MotherEmail  string `gorm:"motherEmail"`
	FatherName   string `gorm:"fatherName"`
	FatherMobile string `gorm:"fatherMobile"`
	FatherEmail  string `gorm:"fatherEmail"`
}