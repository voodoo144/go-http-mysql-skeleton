package main

import (
	"fmt"
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
	log.Println(cfgType)
	for i := 0; i < cfgType.NumField(); i++ {
		field := cfgType.Field(i)
		log.Println(field)
		tag := field.Tag
		envvalue := os.Getenv(string(tag))
		log.Println(envvalue)
		field_value := reflect.ValueOf(&config).Elem().Field(i)
		log.Println(field_value)
		log.Println(field_value.CanSet())
		field_value.SetString(envvalue)
	}
	log.Println(config)
	return &config
}

func main() {
	log.Println("Starting application")
	config := GetServiceConfig()
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("test")
	})
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
