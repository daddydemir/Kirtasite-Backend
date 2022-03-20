package repositories

import (
	"demir/config"
	"demir/models"
)

func FileByUserId(userId string) []models.File {
	var files []models.File
	config.DB.Find(&files, "user_id = ?", userId)
	return files
}

func FileDelete(fileId string) {
	config.DB.Delete(models.File{}, "id = ?", fileId)
}

func FileAdd(file models.File) {
	config.DB.Create(&file)
}
