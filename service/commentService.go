package service

import (
	"github.com/daddydemir/kirtasiye-projesi/auth"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func GetCommentByUserIdService(token string, userId int) (bool, map[string]string) {
	status, message := auth.CommentAuth(token, userId)
	if status {
		return true, map[string]string{}
	} else {
		return false, message
	}
}

func CommentByStationeryIdService(token string) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		return true, map[string]string{}
	} else {
		return false, message
	}
}

func CommentAddService(token string, comment models.Comment) (bool, map[string]string) {
	status, message := auth.CommentAuth(token, comment.UserId)
	if status {
		return true, map[string]string{}
	} else {
		return false, message
	}
}
