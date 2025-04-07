package main

import (
	"log"
	"net/http"

	"github.com/MasterEvarior/gize/cmd/helper"
	"github.com/MasterEvarior/gize/cmd/view"
)

func main() {
	port := helper.GetEnvVarWithDefault("GIZE_PORT", ":8080")

	http.HandleFunc("/", view.Overview)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Could not start the server because of the following issue: %v", err)
	} else {
		log.Println("Ready to accept requests")
	}
}
