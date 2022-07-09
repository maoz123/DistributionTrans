package main

import(
	"fmt"
)

func InsertEvent(order Order) bool {
	if !isInit() {
		Init()
	}
	InsertLocalEvent(order)
	fmt.Println("insert local event succeed")
	return true
}