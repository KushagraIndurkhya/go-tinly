package utills

import (
	"net/http"
)

func Make_Cookie(id string) (*http.Cookie, error) {

	jwt_token, err := CreateToken(id)
	if err != nil {
		return nil, err
	}

	return &(http.Cookie{Name: "Token", Value: jwt_token, MaxAge: 0, HttpOnly: true}), err
}
