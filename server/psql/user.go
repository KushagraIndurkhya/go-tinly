package psql

import (
	"errors"

	"github.com/KushagraIndurkhya/go-tinly/models"
	"github.com/jinzhu/gorm"
)

func Save(u *models.UserInfo) error {
	//TODO:Check validity of data in json
	tup := models.User{Id: u.Id, Email: u.Email, Name: u.Name, Limit: 1000}
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
