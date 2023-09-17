package basic

import (
	"encoding/json"
	"fmt"
	"testing"
)


type User struct{
	FullName string `json:"Name"`
	Age		 int 
	Status string `json:"status"`
}

func TestJson(t *testing.T) {
	//merubah data json ke byte
	var jsonString = `{"Name" : "john wick","age": 27,"status":"Amba"}`
	var jsonData =[]byte(jsonString)
	//menampung json
	var data User
	var err = json.Unmarshal(jsonData,&data)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user : ", data.FullName)
	fmt.Println("age : ", data.Age)
	fmt.Println("status :", data.Status)
}