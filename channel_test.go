package basic

import (
	"fmt"
	"testing"
)


func TestChannel(t *testing.T) {
	var messages = make(chan string)

	var sayHello = func (who string)  {
		var data = fmt.Sprintf("Hello %s", who)
		messages <- data
	}
	go sayHello("Amba")
	go sayHello("Tukam")
	go sayHello("SiImut")

	var messages1 = <-messages
	fmt.Println(messages1)
	var messages2 = <-messages
	fmt.Println(messages2)
	var messages3 = <-messages
	fmt.Println(messages3)


}