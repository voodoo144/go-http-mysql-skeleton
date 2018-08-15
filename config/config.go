package config

import (
	"log"
	"reflect"
	"os"
)

var _config *ServiceConfig

type ServiceConfig struct {
	Port          string "TEST_PORT"
	MysqlUser     string "MYSQL_USER"
	MysqlPassword string "MYSQL_PASSWORD"
	MysqlHost     string "MYSQL_HOST"
	MysqlPort     string "MYSQL_PORT"
	MysqlDatabase string "MYSQL_DATABASE"
}

func GetServiceConfig() *ServiceConfig {
	if _config==nil {
		config := ServiceConfig{}
		log.Println("Trying to get configuration from env")
		cfgType := reflect.TypeOf(config)
		for i := 0; i < cfgType.NumField(); i++ {
			field := cfgType.Field(i)
			tag := field.Tag
			if len(tag) == 0 {
				log.Fatal("Dont you forget to add tag for field: " + field.Name + " ?")
			}
			log.Println("Getting ENV variable: " + tag)
			value := os.Getenv(string(tag))
			log.Println("Value: " + value)
			field_value := reflect.ValueOf(&config).Elem().Field(i)
			field_value.SetString(value)
		}
		log.Println("Final configuration:")
		_config=&config
		log.Println(&config)
	}else{
		log.Println("Already configured")
	}
	return _config
}
