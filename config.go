package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func failedConnectToDB(value interface{})  {
	if value != nil {
		log.Println(value)
		panic("failed to connect database")
	}
}

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "sammidev:sammidev@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)
	failedConnectToDB(err)

	db.Exec("CREATE DATABASE orders_db")
	db.Exec("USE orders_db")
	db.AutoMigrate(&Order{}, &Item{})
}
