package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

// buradaki veriler public olmalı -->

// stationery kısmında adresId alınacak
func GetAddressById(id string) models.Address {
	var address models.Address
	config.DB.Find(&address, "id = ?", id)
	return address
}

func GetAddressByCity(id string) []models.Address {
	var address []models.Address
	config.DB.Find(&address, "city_id = ?", id)
	return address
}

func GetAddressByDistrict(id string) []models.Address {
	var address []models.Address
	config.DB.Find(&address, "district_id = ?", id)
	return address
}
