package psql

import (
	"log"

	"github.com/KushagraIndurkhya/go-tinly/models"
)

func Add_URL(id string, url string) error {

	tup := models.User_URL{Short: url, Created_by: id}
	err := PSQL_DB.Debug().Create(&tup).Error
	if err != nil {
		return err
	}
	return nil
}

func Get_URLs(id string) ([]string, error) {
	urls := []string{}
	// result := PSQL_DB.Select("Short").Where("Created_by = ?", id).Find(&urls)
	// result := PSQL_DB.Raw("SELECT short FROM user_urls WHERE created_by=?", id).Scan(&urls)
	rows, err := PSQL_DB.Raw("SELECT short FROM user_urls WHERE created_by=?", id).Rows()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var short string
		rows.Scan(&short)
		urls = append(urls, short)
	}
	// fmt.Print(urls)
	return urls, nil
}

func Del_URL(sh string) {
	PSQL_DB.Where("short = ?", sh).Delete(&models.User_URL{})
}
