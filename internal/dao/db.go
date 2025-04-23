package dao

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/test_db")
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("数据库无法访问: %v", err)
	}
}

func GetDB() *sql.DB {
	return db
}
