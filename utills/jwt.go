package utills

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id string) (string, error) {
	claim := jwt.MapClaims{}
	claim["authorized"] = true
	claim["user_id"] = id
	claim["exp"] = time.Now().Add(time.Minute * 60 * 24 * 365).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := at.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil

}

func GetId(s string) (bool, string, error) {

	token, err := jwt.Parse(s, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("APP_SECRET")), nil
	})

	if err != nil {
		err = errors.New("couldn't parse claims")
		return false, "", err
	}
	claims := token.Claims.(jwt.MapClaims)

	return true, claims["user_id"].(string), nil

}
