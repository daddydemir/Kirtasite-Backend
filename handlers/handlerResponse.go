package handlers

// token' ın süresi dolmuş ise
func UnauthorizedMessage() map[string]interface{} {
	return map[string]interface{}{"message": "Oturum süreniz dolmuş lütfen oturum açın."}
}

// token yok ise
func NotLoginMessage() map[string]interface{} {
	return map[string]interface{}{"message": "İçeriğe ulaşabilmek için oturum açın."}
}
