package psql

import (
	"errors"
	"log"

	"github.com/KushagraIndurkhya/go-tinly/models"
	"github.com/jinzhu/gorm"
)

func Save(u *models.UserInfo) error {
	//TODO:Check validity of data in json
	tup := models.User{Id: u.Id, Email: u.Email, Name: u.Name, Limit: 10}
	err := PSQL_DB.Debug().Create(&tup).Error

	if err != nil {
		return err
	}
	return nil
}

func Get(id string) (*models.User, error) {

	ret := new(models.User)
	result := PSQL_DB.First(&ret, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return ret, nil
}

func Add_URL(id string, url string) error {

	tup := models.User_URL{Short: url, Created_by: id, Health: true}
	err := PSQL_DB.Debug().Create(&tup).Error
	if err != nil {
		return err
	}
	return nil
}

func Get_URLs(id string) ([]string, error) {
	var urls []string
	result := PSQL_DB.Find(&urls, "Created_by = ?", id)

	if result.Error != nil {
		log.Print(result.Error)
		return nil, result.Error
	}
	return urls, nil
}
