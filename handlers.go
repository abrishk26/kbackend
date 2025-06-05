package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func healthCheck(w http.ResponseWriter, r *http.Request, _  httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}