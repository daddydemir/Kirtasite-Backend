package repositories

import (
	"demir/config"
	"demir/models"
)

func GetCommentByUserId() []models.Comment {
	var comments []models.Comment
	config.DB.Find(&comments, "user_id = ?", "9")
	return comments
}
