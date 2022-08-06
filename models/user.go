package models

type User struct {
	Id        int    `json:"id"`
	RoleId    int    `json:"role_id"`
	Password  string `json:"-"`
	ImagePath string `json:"image_path"`
	Phone     string `json:"phone"`
	Mail      string `json:"mail"`
}
type Users struct {
	Id        int    `json:"id"`
	RoleId    int    `json:"role_id"`
	Password  string `json:"-"`
	ImagePath string `json:"image_path"`
	Phone     string `json:"phone"`
	Mail      string `json:"mail"`
	RoleData  Role   `json:"role" gorm:"foreignKey:role_id;references:id"`
}
