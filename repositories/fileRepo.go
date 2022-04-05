package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
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

func FileById(id string) models.File {
	var file models.File
	config.DB.Find(&file, "id = ?", id)
	return file
}
