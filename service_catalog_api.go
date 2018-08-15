package main

import (
	"ServiceCatalogApi/handlers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"reflect"
)

type ServiceConfig struct {
	Port          string "TEST_PORT"
	MysqlUser     string "MYSQL_USER"
	MysqlPassword string "MYSQL_PASSWORD"
	MysqlPort     string "MYSQL_HOST"
	MysqlDatabase string "MYSQL_DATABASE"
}

func GetServiceConfig() *ServiceConfig {
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
		envvalue := os.Getenv(string(tag))
		log.Println("Value: " + envvalue)
		field_value := reflect.ValueOf(&config).Elem().Field(i)
		field_value.SetString(envvalue)
	}
	log.Println("Final configuration:")
	log.Println(config)
	return &config
}

func main() {
	log.Println("Starting application")
	config := GetServiceConfig()
	router := httprouter.New()
	router.GET("/", handlers.BasicHandler)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
