package repositories

import (
	"demir/config"
	"demir/models"
)

func FileGetByUserId() []models.File {
	var files []models.File
	config.DB.Find(&files, "user_id = ?", "9")
	return files
}
