package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)
func main() {
	// Router
	router := httprouter.New()

	router.GET("/health_check", healthCheck)

	// Server
	server := &http.Server {
		Addr: ":8080",
		Handler: router,
	}

	log.Println("✅ Server is running at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

