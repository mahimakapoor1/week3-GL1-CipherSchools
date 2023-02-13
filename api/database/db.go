package database

import (
	"fmt"
	"gin/basic/models"
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
func GetDB() *gorm.DB{
	return DB
}
func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	username := "postgres"
	password := "postgres"
	args := "hosts=" + host + " port=" + port + " user=" + username + " dbname=" + dbName + " sslmode=disable password"+password
	db, err := gorm.Open("postgres", args)
	if err!=nil{
		log.Fatal(err)
	}
		db.AutoMigrate(models.Book{})
	DB=db;
}
func exam(){
	fmt.Print("add")
}