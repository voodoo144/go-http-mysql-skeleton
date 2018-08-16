package database

import (
	. "ServiceCatalogApi/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	_db     *gorm.DB
	err     error
	dialect string = "mysql"
)

func OpenConnection() {
	c := GetServiceConfig()
	_db, err := gorm.Open(dialect, c.MysqlUser+":"+c.MysqlPassword+"@/"+c.MysqlDatabase+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
	}
	defer _db.Close()
}

func ExecuteQuery(query string) {
	if _db == nil {
		OpenConnection()
	}
	log.Println("Trying to execute query: " + query)
	tx := _db.Begin()
	if tx.Error != nil {
		log.Println("Error starting transaction:")
		log.Println(tx.Error)
	}
	if err := tx.Exec(query).Error; err != nil {
		log.Println(err)
		log.Println("Error in transaction, rolling back")
		tx.Rollback()
	}
	log.Println("Commiting transaction")
	tx.Commit()
}
