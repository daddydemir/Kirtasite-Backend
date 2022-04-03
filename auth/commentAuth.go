package auth

import "demir/repositories"

func CommentAuth(tokenString string, userId int) (bool, map[string]string) {
	status, message := IsValid(tokenString)
	if status {
		tokenUser, state := repositories.UserByName(TokenParser(tokenString))
		if state {
			if tokenUser.Id == userId {
				return true, map[string]string{}
			} else {
				return false, map[string]string{"message": "Yetksisiz kullanıcı."}
			}
		} else {
			return false, map[string]string{"message": "Yetksisiz kullanıcı."}
		}
	} else {
		return false, message
	}
}
