package models

import "time"

type File struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Private     bool      `json:"private"`
	FilePath    string    `json:"file_path"`
	CreatedDate time.Time `json:"created_date"`
}
