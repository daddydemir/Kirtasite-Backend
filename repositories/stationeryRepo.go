package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/security"
)

func GetStationeryByUserId(stationeryId string) (interface{}, bool) {
	var stationery models.Stationeries
	result := config.DB.Find(&stationery, "user_id = ?", stationeryId)
	return returnModel(result.Error, stationery)
}
func GetStationeryByName(companyName string) (interface{}, bool) {
	var stationery models.Stationeries
	result := config.DB.Find(&stationery, "company_name = ?", companyName)
	return returnModel(result.Error, stationery)
}
func AddStationery(stationery models.Stationeries) (interface{}, bool) {
	_, status := GetStationeryByName(stationery.CompanyName)
	if status {
		return "Bu isimde Kayıtlı bir kırtasiye zaten var.", false
	} else {
		stationery.UserData.Password = security.HashPassword(stationery.UserData.Password)
		result := config.DB.Create(&stationery)
		return returnModel(result.Error, stationery)
	}

}
func DeleteStationery(stationeryId string) (interface{}, bool) {
	result := config.DB.Delete(models.Stationery{}, "id = ?", stationeryId)
	return returnModel(result.Error, "")
}
func UpdateStationery(stationery models.Stationeries) (interface{}, bool) {
	// TODO parola sıfırlama isteğini kontrol et
	result := config.DB.Save(&stationery)
	return returnModel(result.Error, stationery)
}
func AllStationery() (interface{}, bool) {
	var stationery []models.Stationeries
	result := config.DB.Find(&stationery)
	return returnModel(result.Error, stationery)
}
func UpdateStationeryImage(url string, stationeryId string) (interface{}, bool) {
	response, _ := GetStationeryByUserId(stationeryId)
	var stationery models.Stationeries
	stationery = response.(models.Stationeries)
	stationery.UserData.ImagePath = url
	result := config.DB.Save(&stationery)
	return returnModel(result.Error, stationery)
}
