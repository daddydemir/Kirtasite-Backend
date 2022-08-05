package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func GetFileByCustomerId(userId string) (interface{}, bool) {
	var files []models.File
	result := config.DB.Find(&files, "customer_id = ?", userId)
	return returnModel(result.Error, files)
}
func FileDelete(fileId string) {
	config.DB.Delete(models.File{}, "id = ?", fileId)
}
func AddFile(file models.File) (interface{}, bool) {
	result := config.DB.Create(&file)
	return returnModel(result.Error, file)
}
func GetFileById(id string) (interface{}, bool) {
	var file models.Files
	result := config.DB.Find(&file, "id = ?", id)
	return returnModel(result.Error, file)
}
