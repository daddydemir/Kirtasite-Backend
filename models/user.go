package models

type User struct {
	Id       int    `json:"id"`
	RoleId   int    `json:"role_id" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
	TelNo    string `json:"tel_no" binding:"required"`
	//UserRole Role   `gorm:"foreignKey:role_id"`
}
