package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	. "ServiceCatalogApi/config"
)

var db *gorm.DB
var err error
var dialect string = "mysql"

func OpenConnection() {
	c:=GetServiceConfig()
	db, err := gorm.Open(dialect, c.MysqlUser+":"+c.MysqlPassword+"@/"+c.MysqlDatabase+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
}

func ExecuteQuery(query string) {
	if db == nil {
		OpenConnection()
	}
	tx := db.Begin()
	if tx.Error != nil {
		log.Println("Error starting transaction:")
		log.Println(tx.Error)
	}
	if err := tx.Exec(query).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
}
