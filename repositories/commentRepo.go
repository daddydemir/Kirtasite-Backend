package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func GetCommentByCustomerId(userId string) (interface{}, bool) {
	var comments []models.Comments
	result := config.DB.Find(&comments, "customer_id = ?", userId)
	return returnModel(result.Error, comments)
}
func GetCommentByStationeryId(stationeryId string) (interface{}, bool) {
	var comments []models.Comments
	result := config.DB.Find(&comments, "stationery_id = ?", stationeryId)
	return returnModel(result.Error, comments)
}
func AddComment(comment models.Comment) (interface{}, bool) {
	result := config.DB.Create(&comment)
	return returnModel(result.Error, comment)
}
