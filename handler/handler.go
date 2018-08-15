package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"ServiceCatalogApi/database"
)

func BasicHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	database.ExecuteQuery("Select * from test")
}
