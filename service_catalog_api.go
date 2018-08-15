package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

type ServiceConfig struct {
	Port          string
	MysqlUser     string
	MysqlPort     string
	MysqlDatabase string
}

func GetServiceConfig() *ServiceConfig {
	config := ServiceConfig{
		Port:          os.Getenv("TEST_PORT"),
		MysqlUser:     os.Getenv(""),
		MysqlPort:     os.Getenv(""),
		MysqlDatabase: os.Getenv(""),
	}

	return &config
}

func main() {
	config := GetServiceConfig()
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("test")
	})
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
