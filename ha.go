package main

import (
	"database/sql"


	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"unicode/utf8"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:xingfushenghuo999@tcp(127.0.0.1:3306)/college_journal?collation=utf8mb4_unicode_ci&charset=utf8mb4&timeout=5s&readTimeout=15s&writeTimeout=15s")
	if err != nil {
		panic(err)
	}
	fmt.Println(db.Ping())
	db_user_p, _ := db.Prepare("INSERT INTO user VALUES(?,?,?,?,?,?,now(),?,?,?)")

	s, err := db_user_p.Exec(0, "18208142443", "大托普", utf8.ValidString("恒少"), utf8.ValidString("是人"), utf8.ValidString("Sd"), "xingfusehnghuoa", "normal", "2")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	fmt.Println(utf8.ValidString("恒少"))


}