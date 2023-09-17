package basic

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParsingUrl(t *testing.T) {
	var urlString = "http://dreammybull.com:80/hello?name=john wick&age=27"
	var u,e = url.Parse(urlString)//parsing string ke url
	if e != nil{
		fmt.Println(e.Error())
		return
	}
	fmt.Printf("Url: %s\n", urlString)

	fmt.Printf("Protokol: %s\n", u.Scheme)
	fmt.Printf("Host: %s\n", u.Host)
	fmt.Printf("Path: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name : %s, age: %s",name,age)
	

}