package models

type User struct {
	Id        int    `json:"id"`
	RoleId    int    `json:"role_id"`
	Username  string `json:"username" validate:"required,min=5"`
	Password  string `json:"password" validate:"required,min=8"`
	ImagePath string `json:"image_path" validate:"url"`
	Mail      string `json:"mail" validate:"required,email"`
	TelNo     string `json:"tel_no" validate:"required,len=10"`
	//UserRole Role   `gorm:"foreignKey:role_id"`
}
