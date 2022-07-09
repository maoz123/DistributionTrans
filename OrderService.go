package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	OrderId int

	Money int

	Amount int

	Desc string
}

func ReductStorage(order Order) bool {
	if !isInit() {
		Init()
	}
	BeginTransaction()
	UpdateStorage(order.Amount)
	InsertEvent(order)
	CommitTransaction()
	fmt.Printf("distribution transaction first step completed! %d \n", order.OrderId)

	var client = GetClient()
	var message, _ = json.Marshal(order)
	SendMessage(string(message), client)
	// If transaction completed, we need to send events to service bus queue.
	fmt.Printf("send message completed %d!\n", order.OrderId)
	return true
}
