package psql

import (
	"fmt"
	"log"

	"github.com/KushagraIndurkhya/go-tinly/models"
)

func Add_URL(id string, url string,comments string,medium string,source string,campaign string,keyword string ) error {

	tup := models.User_URL{Short: url, Created_by: id, Comments: comments, Medium: medium, Source: source, Campaign: campaign, Keyword: keyword}
	err := PSQL_DB.Debug().Create(&tup).Error
	if err != nil {
		return err
	}
	return nil
}

func Get_URLs(id string) ([]models.URL_INFO_RESPONSE, error) {
	urls := []models.URL_INFO_RESPONSE{}
	// result := PSQL_DB.Select("Short").Where("Created_by = ?", id).Find(&urls)
	// result := PSQL_DB.Raw("SELECT short FROM user_urls WHERE created_by=?", id).Scan(&urls)
	rows, err := PSQL_DB.Raw("SELECT short,Comments,Medium,Source,Campaign,Keyword FROM user_urls WHERE created_by=?", id).Rows()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		fmt.Print(rows)
		var short string
		var comments string
		var medium string
		var source string
		var campaign string
		var keyword string
		err := rows.Scan(&short, &comments, &medium, &source, &campaign, &keyword)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		urls = append(urls, models.URL_INFO_RESPONSE{Short: short, Comments: comments, Medium: medium, Source: source, Campaign: campaign, Keyword: keyword})
	}
	// fmt.Print(urls)
	return urls, nil
}

func Del_URL(sh string) error {
	//Return Error
	PSQL_DB.Where("short = ?", sh).Unscoped().Delete(&models.User_URL{})
	return nil
}
