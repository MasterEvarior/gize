package main

import (
	"log"
	"net/http"

	"github.com/MasterEvarior/gize/cmd/view"
)

func main() {

	http.HandleFunc("/", view.Overview)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Could not start the server because of the following issue: %v", err)
	} else {
		log.Println("Ready to accept requests")
	}
}
