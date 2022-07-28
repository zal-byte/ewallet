package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"
const DB_NAME = "ewallet"

func DbConn() *gorm.DB {

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, err := xorm.NewEngine("mysql", dsn)

	// if error != nil {
	// 	panic(error)
	// }

	if err != nil {
		panic(err)
	}
	return db
}
