package main

import (
	"ServiceCatalogApi/handler"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	. "ServiceCatalogApi/config"
)


func main() {
	log.Println("Starting application")
	config := GetServiceConfig()
	router := httprouter.New()
	router.GET("/", handler.BasicHandler)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
