package main

import (
	"fmt"
	//"encoding/json"
)

func main(){
	client := GetClient()
	/*
	message := &Order{
		OrderId : 1,
		Money : 10,
		Desc : "test order",
		Amount : 3,
	}
	brokerMessage,_ := json.Marshal(message);
	SendMessage(string(brokerMessage), client)
	*/
	GetMessage(1, client)
	fmt.Println("send message")
}