package redis_layer

import (
	"encoding/json"
	"fmt"
)

type Url_req struct {
	Url        string
	Created_by string
	Created_at int64
}

func Get_Val(u *Url_req) (string, error) {
	res, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error in encoding")
		return "", err
	}
	return string(res), nil
}
