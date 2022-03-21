package models

import "time"

type File struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id" validate:"required"`
	Private     bool      `json:"private" validate:"required"`
	FilePath    string    `json:"file_path" validate:"required,url"`
	CreatedDate time.Time `json:"created_date"`
}
