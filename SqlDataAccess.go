package main

import (
	"database/sql"
	"fmt"
	"net/url"
)

var db *sql.DB
var initStatus bool = false
var tx *sql.Tx

type Params struct {
	UserName string

	Password string

	HostName string

	Port int
}

var params = &Params{
	UserName: "",

	Password: "",

	HostName: "",

	Port: 10,
}

func isInit() bool {
	return initStatus
}

/*
	init method must be called first.
*/
func Init() {
	query := url.Values{}
	query.Add("app name", "MyAppName")
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(params.UserName, params.Password),
		Host:   fmt.Sprintf("%s:%d", params.HostName, params.Port),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	db, _ = sql.Open("sqlserver", u.String())
	initStatus = true
}

func BeginTransaction() {
	tx, _ = db.Begin()
}

func CommitTransaction() {
	tx.Commit()
}

func UpdateStorage(amount int) int64 {
	var script = "UPDATE Storage Set Amount = Amount - ?"

	result, _ := db.Exec(script, amount)

	want, _ := result.RowsAffected()
	fmt.Println("update succeed!")
	return want
}

func InsertLocalEvent(order Order) bool {
	var script = "INSERT INTO Storage (OrderId, Money, Amount, Desc) VALUES (?, ? ,? ?)"

	result, _ := db.Exec(script, order.OrderId, order.Money, order.Amount, order.Desc)

	wanted, _ := result.RowsAffected()
	fmt.Println("insert succeed!")
	if wanted > 0 {
		return true
	}
	return false
}
