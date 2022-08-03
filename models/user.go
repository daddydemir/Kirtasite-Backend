package models

type User struct {
	Id        int
	RoleId    int
	Password  string
	ImagePath string
	Phone     string
	Mail      string
}
