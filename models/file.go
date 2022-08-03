package models

import "time"

type File struct {
	Id          int
	UserId      int
	Private     bool
	FilePath    string
	FolderId    string
	CreatedDate time.Time
}
