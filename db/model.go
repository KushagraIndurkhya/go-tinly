package db

import(
	"fmt"
    "encoding/json"
)

type Url_req struct{
	Url string
	Count int
	Created_by string
}

func Get_Val(u *Url_req) (string,error){
	res,err:=json.Marshal(u)
	if err!=nil{
		fmt.Println("error in encoding")
		return "",err	
	}
	return string(res),nil
}
