package basic

import (
	"encoding/base64"
	"fmt"
	"testing"
)


func TestEncode(t *testing.T) {
	var data = "Ambatukam"
	//Encode
	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)
	//Decode
	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("decode :", decodedString)

	var data2 = "https://ambatukam.id"
	
	var encodedURL = base64.URLEncoding.EncodeToString([]byte(data2))
	fmt.Println(encodedURL)

	var decodedByteUrl, _ = base64.URLEncoding.DecodeString(encodedURL)
	var decodedStringUrl = string(decodedByteUrl)
	fmt.Println(decodedStringUrl)

}