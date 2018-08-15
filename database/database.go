package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB
var err error

type ConnectionParams struct {
	Host     string
	Port     string
	User     string
	Password string
}

func InitDatabase(host string, port string, user string, password string) {

}

func OpenConnection() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func ExecuteQuery(query string) {
	if db == nil {
		OpenConnection()
	}
	tx := db.Begin()
	if tx.Error != nil {
		//
	}
	if err := tx.Exec(query).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()

}
