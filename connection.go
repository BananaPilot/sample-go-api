package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connection(uri string) *sql.DB {
	conn, err := sql.Open("mysql", uri)
	if err != nil {
		panic(err.Error())
	}
	return conn
}

func CreateUri(user, password, ip, port, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, ip, port, database)
}
