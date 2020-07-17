package main

import (
	"github.com/duniapay/api/http/router"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.WithFields(log.Fields{
		"Service": "messaging-service",
	}).Info("STARTING SERVER ON PORT 80")

	// Start router
	newRouter := router.New()

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":80", handlers.CORS()(newRouter)))
}
