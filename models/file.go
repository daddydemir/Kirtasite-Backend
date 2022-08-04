package models

import "time"

type File struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Private     bool      `json:"private"`
	FilePath    string    `json:"file_path"`
	FolderId    string    `json:"folder_id"`
	CreatedDate time.Time `json:"created_date"`
}
