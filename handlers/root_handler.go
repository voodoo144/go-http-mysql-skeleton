package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func BasicHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("test")
}
