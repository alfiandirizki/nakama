package basic

import (
	"fmt"
	"os"
	"testing"
)


var path = "/home/ylen/back-end/basic/amba.txt"

func isError(err error)bool  {
	if err != nil{
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func CreateFile()  {
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)

	if os.IsNotExist(err){
		var file,err = os.Create(path)
		if isError(err) { return}
		defer file.Close()
	}
	fmt.Println("file berhasil dibuat",path)

}

func Test(t *testing.T) {
	CreateFile()
}