package models

type Stationery struct {
	Id          int     `json:"id"`
	RoleId      int     `json:"role_id" validate:"required"`
	CompanyName string  `json:"company_name" validate:"required,min=5"`
	Password    string  `json:"password" validate:"required,min=8"`
	ImagePath   string  `json:"image_path" validate:"required,url"`
	Mail        string  `json:"mail" validate:"required,email"`
	TelNo       string  `json:"tel_no" validate:"required,len=9"`
	AddressId   string  `json:"address_id" validate:"required"`
	Score       float32 `json:"score" validate:"required,gt=0,lt=11"`
}
