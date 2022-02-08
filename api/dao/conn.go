package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:Andy401927!@tcp(127.0.0.1:3306)/video_server?charset=utf8")
	if err != nil {
		log.Printf("Connect db failed:%s", err)
		panic(err)
	}
}
