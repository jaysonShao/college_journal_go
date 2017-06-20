package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "name:password@tcp(127.0.0.1:3306)/college_journal?timeout=90s&collation=utf8mb4_unicode_ci")
	if err != nil {
		panic(err)
	}
	fmt.Println(db.Ping())
}

func CloseDb(){
	defer db.Close()
}