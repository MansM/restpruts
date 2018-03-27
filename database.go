package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//NewDB creates  new DB connection,
//TODO: should get values from env
func NewDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=db port=26257 user=gorm dbname=events password=pass sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	db.AutoMigrate(&Event{})
	//defer db.Close()

	return db
}
