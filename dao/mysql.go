package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	var dns = "root:123456@tcp(127.0.0.1:3306)/rpc?charset=utf8mb4&parseTime=True&loc=Local" //链接数据库
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10) //设置最大的连接数
	db.SetMaxIdleConns(5)  //设置最大空闲数

	err = db.Ping() //检查数据库链接问题
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	return db, nil
}
