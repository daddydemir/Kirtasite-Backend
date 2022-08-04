package models

type User struct {
	Id        int    `json:"id"`
	RoleId    int    `json:"role_id"`
	Password  string `json:"-"`
	ImagePath string `json:"image_path"`
	Phone     string `json:"phone"`
	Mail      string `json:"mail"`
}
