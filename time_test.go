package basic

import (
	"fmt"
	"testing"
	"time"
)



func TestTime(t *testing.T) {
	//Default 
	var time1 = time.Now()
	fmt.Printf("Waktu : %v\n", time1)
	//Kita tentukan/custom
	var time2 = time.Date(2021,12,24,10,20,0,0,time.UTC)
	fmt.Printf("time2 %v\n", time2)

}