package models

type Stationery struct {
	Id          int     `json:"id"`
	RoleId      int     `json:"role_id"`
	CompanyName string  `json:"company_name"`
	Password    string  `json:"password"`
	ImagePath   string  `json:"image_path"`
	Mail        string  `json:"mail"`
	TelNo       string  `json:"tel_no"`
	Address     string  `json:"address"`
	Score       float32 `json:"score"`
}
