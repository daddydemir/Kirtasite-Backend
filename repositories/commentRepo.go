package repositories

import (
	"demir/config"
	"demir/models"
)

func CommentByUserId(userId string) []models.Comment {
	var comments []models.Comment
	config.DB.Find(&comments, "user_id = ?", userId)
	return comments
}

func CommentByStationeryId(stationeryId string) []models.Comment {
	var comments []models.Comment
	config.DB.Find(&comments, "stationery_id = ?", stationeryId)
	return comments
}

func CommentAdd(comment models.Comment) {
	config.DB.Create(&comment)
}
