package models

type Student struct {
	ID             int    `gorm:"primaryKey,id"`
	Password       string `gorm:"password"`
	IdCardNumber   string `gorm:"id_card_number"`
	Name           string `gorm:"name"`
	ClassId        string `gorm:"class_id"`
	Gender         string `gorm:"gender"`
	Dob            string `gorm:"dob"`
	Mobile         string `gorm:"mobile"`
	ParentsInfoId  int    `gorm:"parents_info_id"`
	Address        string `gorm:"address"`
	BloodGroup     string `gorm:"blood_group"`
	SecondLanguage string `gorm:"second_language"`
	ProfilePicLink string `gorm:"profile_pic_Link"`
	Remarks        string `gorm:"remarks"`
	Religion       string `gorm:"religion"`
	Caste          string `gorm:"caste"`
	Pwd            bool   `gorm:"pwd"`
	Nationality    string `gorm:"nationality"`
}

type Staff struct {
	ID              int    `gorm:"primaryKey,id"`
	Name            string `gorm:"name"`
	HashedPassword  string `gorm:"hashed_password"`
	Address         string `gorm:"address"`
	Mobile          string `gorm:"mobile"`
	IsTeachingStaff bool   `gorm:"is_teaching_staff"`
	AdditionalRole  string `gorm:"additional_role"`
}

type Admin struct {
	ID             int    `gorm:"primaryKey,id"`
	HashedPassword string `gorm:"hashedpassword"`
	Email          string `gorm:"email"`
	Name           string `gorm:"name"`
	Mobile         string `gorm:"mobile"`
	Address        string `gorm:"address"`
}

type ParentInfo struct {
	ID           int    `gorm:"primaryKey,id"`
	MotherName   string `gorm:"mother_name"`
	MotherMobile string `gorm:"mother_mobile"`
	MotherEmail  string `gorm:"mother_email"`
	FatherName   string `gorm:"father_name"`
	FatherMobile string `gorm:"father_mobile"`
	FatherEmail  string `gorm:"father_email"`
	GuardianName string `gorm:"guardian_name"`
	GuardianMobile string `gorm:"guardian_mobile"`
	GuardianEmail  string `gorm:"guardian_email"`
}