package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)


type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func connect() {
	// Connect to the database
	db, err := gorm.Open(mysql.Open("naol:naol.@/ForTest?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		os.Exit(2)
	}

	log.Println("Database connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migration")


	Database = DbInstance{Db: db}

	Database = 

}