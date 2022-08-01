package database

import (
	"ewallet/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"
const DB_NAME = "ewallet"

func DbConn() *gorm.DB {

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Wallets{})
	if err != nil {
		panic(err)
	}
	return db
}

func XormConn() *xorm.Engine {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME

	db, err := xorm.NewEngine("mysql", dsn)

	if err != nil {
		panic(err)
	}
	return db

}
